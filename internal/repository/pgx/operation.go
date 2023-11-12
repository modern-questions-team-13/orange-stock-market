package pgx

import (
	"context"
	"fmt"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
)

type Operation struct {
	pg *database.Postgres
}

func (o *Operation) Create(ctx context.Context, buyerId, sellerId, companyId int, price int) error {
	sql, args, err := o.pg.Sq.Insert("operations").
		Columns("buyer_id", "seller_id", "company_id", "price").
		Values(buyerId, sellerId, companyId, price).ToSql()

	if err != nil {
		return err
	}

	res, err := o.pg.Pool.Exec(ctx, sql, args...)

	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("error store operation")
	}

	return nil
}

func NewOperation(pg *database.Postgres) *Operation {
	return &Operation{pg: pg}
}
