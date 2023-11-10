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

type Buy struct {
	pg *database.Postgres
}

func (b *Buy) Get(ctx context.Context, id int) (model.Buy, error) {
	sql, args, err := b.pg.Sq.Select("buys").
		Columns("user_id", "company_id", "price", "created_at").
		Where("id = $", id).
		ToSql()

	if err != nil {
		return model.Buy{}, err
	}

	var (
		userId    int
		companyId int
		price     int
		createdAt time.Time
	)

	err = b.pg.Pool.QueryRow(ctx, sql, args...).Scan(&userId, &companyId, &price, &createdAt)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				return model.Buy{}, fmt.Errorf("error get buy with id=%d: %w", id, repoerrs.ErrNotExists)
			}
		}

		return model.Buy{}, err
	}

	return model.Buy{Id: id, UserId: userId, CompanyId: companyId, Price: price, CreatedAt: createdAt}, err
}

func (b *Buy) GetBuys(ctx context.Context, companyId int, maxPrice int, limit uint64) (id []int, err error) {
	sql, args, err := b.pg.Sq.Select("buys").
		Columns("id").
		Where("company_id=$ and price <= $", companyId, maxPrice).
		Limit(limit).ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := b.pg.Pool.Query(ctx, sql, args...)

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

func (b *Buy) GetAllBuys(ctx context.Context, companyId int, limit, offset uint64) (price []int, err error) {
	sql, args, err := b.pg.Sq.Select("buys").
		Columns("price").
		Where("company_id=$", companyId).
		Limit(limit).Offset(offset).ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := b.pg.Pool.Query(ctx, sql, args...)

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

func NewBuy(pg *database.Postgres) *Buy {
	return &Buy{pg: pg}
}

func (b *Buy) Create(ctx context.Context, userId, companyId int, price int) (model.Buy, error) {
	sql, args, err := b.pg.Sq.Insert("buys").
		Columns("user_id", "company_id", "price").
		Values(userId, companyId, price).
		Suffix("returning id, created_at").
		ToSql()

	if err != nil {
		return model.Buy{}, err
	}

	var (
		id        int
		createdAt time.Time
	)

	err = b.pg.Pool.QueryRow(ctx, sql, args...).Scan(&id, &createdAt)

	if err != nil {
		return model.Buy{}, err
	}

	return model.Buy{Id: id, UserId: userId, CompanyId: companyId, CreatedAt: createdAt, Price: price}, nil

}

func (b *Buy) Delete(ctx context.Context, id int) error {
	buy, err := b.Get(ctx, id)

	if err != nil {
		return err
	}

	sql, args, err := b.pg.Sq.Delete("sales").Where("id = $", id).ToSql()

	if err != nil {
		return err
	}

	tx, err := b.pg.Pool.Begin(ctx)

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

	sql, args, err = b.pg.Sq.Update("users").
		Set("wealth", fmt.Sprintf("wealth + %d", buy.Price)).
		Where("id = $", buy.UserId).ToSql()

	if err != nil {
		return err
	}

	res, err = tx.Exec(ctx, sql, args...)

	if res.RowsAffected() == 0 {
		return fmt.Errorf("error updating balance for user=%d", buy.UserId)
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
