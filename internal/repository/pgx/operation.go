package pgx

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
)

type Operation struct {
	pg *database.Postgres
}

func NewOperation(pg *database.Postgres) *Operation {
	return &Operation{pg: pg}
}

func (o *Operation) Create(ctx context.Context, userId, companyId int, price int) (model.Operation, error) {
	//TODO implement me
	panic("implement me")
}
