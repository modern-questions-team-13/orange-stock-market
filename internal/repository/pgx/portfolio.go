package pgx

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository/repoerrs"
)

type Portfolio struct {
	pg *database.Postgres
}

func (p *Portfolio) Get(ctx context.Context, userId int) ([]model.Asset, error) {
	sql, args, err := p.pg.Sq.Select("company_id", "count").
		From("stock_portfolios").Where("user_id = ? and count > 0", userId).ToSql()

	rows, err := p.pg.Pool.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	assets := make([]model.Asset, 0)

	for rows.Next() {
		var asset model.Asset

		err = rows.Scan(&asset.CompanyId, &asset.Count)

		if err != nil {
			return nil, err
		}

		assets = append(assets, asset)
	}

	return assets, nil
}

func (p *Portfolio) Create(ctx context.Context, userId, companyId int, count int) error {
	sql, args, err := p.pg.Sq.Insert("stock_portfolios").
		Columns("user_id", "company_id", "count").
		Values(userId, companyId, count).
		Suffix("returning id").
		ToSql()

	if err != nil {
		return err
	}

	var id int

	err = p.pg.Pool.QueryRow(ctx, sql, args...).Scan(&id)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return repoerrs.ErrAlreadyExists
			}
		}

		return err
	}

	return nil
}

func (p *Portfolio) AddStock(ctx context.Context, userId, companyId int) error {
	sql := `update stock_portfolios set count=count+1 where user_id=$1 and company_id=$2`

	_, err := p.pg.Pool.Exec(ctx, sql, userId, companyId)

	return err
}

func (p *Portfolio) RemoveStock(ctx context.Context, userId, companyId int) error {
	sql := `update stock_portfolios set count=count-1 where user_id=$1 and company_id=$2`

	res, err := p.pg.Pool.Exec(ctx, sql, userId, companyId)

	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("stock not removed from portfolio: %w", repoerrs.ErrNotExists)
	}

	return nil
}

func NewPortfolio(pg *database.Postgres) *Portfolio {
	return &Portfolio{pg: pg}
}
