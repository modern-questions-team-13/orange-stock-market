package pgx

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository/repoerrs"
	"time"
)

type Sale struct {
	pg *database.Postgres
}

func (s *Sale) GetAllSales(ctx context.Context, companyId int, limit, offset uint64) (price []int, err error) {
	sql, args, err := s.pg.Sq.Select("sales").
		Columns("price").
		Where("company_id=$", companyId).
		Limit(limit).Offset(offset).ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := s.pg.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	price = make([]int, 0, rows.CommandTag().RowsAffected())

	for rows.Next() {
		var curPrice int

		err := rows.Scan(&curPrice)

		if err != nil {
			return nil, err
		}

		price = append(price, curPrice)
	}

	return price, nil
}

func (s *Sale) Get(ctx context.Context, id int) (model.Sale, error) {
	sql, args, err := s.pg.Sq.Select("sales").
		Columns("user_id", "company_id", "price", "created_at").
		Where("id = $", id).
		ToSql()

	if err != nil {
		return model.Sale{}, err
	}

	var (
		userId    int
		companyId int
		price     int
		createdAt time.Time
	)

	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(&userId, &companyId, &price, &createdAt)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				return model.Sale{}, fmt.Errorf("error get sale with id=%d: %w", id, repoerrs.ErrNotExists)
			}
		}

		return model.Sale{}, err
	}

	return model.Sale{Id: id, UserId: userId, CompanyId: companyId, Price: price, CreatedAt: createdAt}, err

}

func (s *Sale) GetSales(ctx context.Context, companyId int, maxPrice int) (id []int, err error) {
	sql, args, err := s.pg.Sq.Select("sales").
		Columns("id").
		Where("company_id=$ and price <= $", companyId, maxPrice).ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := s.pg.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	ids := make([]int, 0, rows.CommandTag().RowsAffected())

	for rows.Next() {
		var curId int

		err := rows.Scan(&curId)

		if err != nil {
			return nil, err
		}

		ids = append(ids, curId)
	}

	return ids, nil
}

func NewSale(pg *database.Postgres) *Sale {
	return &Sale{pg: pg}
}

func (s *Sale) Create(ctx context.Context, userId, companyId int, price int) (model.Sale, error) {
	sql, args, err := s.pg.Sq.Insert("sales").
		Columns("user_id", "company_id", "price").
		Values(userId, companyId, price).
		Suffix("returning id, created_at").
		ToSql()

	if err != nil {
		return model.Sale{}, err
	}

	var (
		id        int
		createdAt time.Time
	)

	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(&id, &createdAt)

	if err != nil {
		return model.Sale{}, err
	}

	return model.Sale{Id: id, UserId: userId, CompanyId: companyId, CreatedAt: createdAt, Price: price}, nil
}

func (s *Sale) Delete(ctx context.Context, id int) error {
	sale, err := s.Get(ctx, id)

	if err != nil {
		return err
	}

	sql, args, err := s.pg.Sq.Delete("sales").Where("id = $", id).ToSql()

	if err != nil {
		return err
	}

	tx, err := s.pg.Pool.Begin(ctx)

	if err != nil {
		return err
	}

	res, err := tx.Exec(ctx, sql, args...)

	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	if res.RowsAffected() == 0 {
		_ = tx.Rollback(ctx)
		return repoerrs.ErrNotExists
	}

	sql, args, err = s.pg.Sq.Update("users").
		Set("balance", fmt.Sprintf("balance + %d", sale.Price)).
		Where("id = $", sale.UserId).ToSql()

	if err != nil {
		return err
	}

	res, err = tx.Exec(ctx, sql, args...)

	if res.RowsAffected() == 0 {
		return fmt.Errorf("error updateing balance for user=%d", sale.UserId)
	}

	if err != nil {
		return err
	}

	err = tx.Commit(ctx)

	if err != nil {
		return fmt.Errorf("error delete sale=%d: %w", id, err)
	}

	return nil
}
