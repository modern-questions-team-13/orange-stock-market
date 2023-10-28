package pgx

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
)

type Buy struct {
	pg *database.Postgres
}

func NewBuy(pg *database.Postgres) *Buy {
	return &Buy{pg: pg}
}

func (b *Buy) Create(ctx context.Context, userId, companyId int, price int) (model.Buy, error) {
	//TODO implement me
	panic("implement me")
}

func (b *Buy) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
