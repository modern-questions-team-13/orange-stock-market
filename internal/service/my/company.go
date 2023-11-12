package my

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
)

type Company struct {
	repos *repository.Repositories
}

func (c *Company) GetAll(ctx context.Context) ([]model.Company, error) {
	return c.repos.Company.GetAll(ctx)
}

func NewCompany(repos *repository.Repositories) *Company {
	return &Company{repos: repos}
}
