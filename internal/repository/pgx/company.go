package pgx

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
)

type Company struct {
	pg *database.Postgres
}

func NewCompany(pg *database.Postgres) *Company {
	return &Company{pg: pg}
}

func (c *Company) Create(ctx context.Context, name string) (model.Company, error) {
	//TODO implement me
	panic("implement me")
}
