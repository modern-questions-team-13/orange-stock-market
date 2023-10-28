package pgx

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
)

type Sale struct {
	pg *database.Postgres
}

func NewSale(pg *database.Postgres) *Sale {
	return &Sale{pg: pg}
}

func (s *Sale) Create(ctx context.Context, userId, companyId int, price int) (model.Sale, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Sale) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
