package my

import (
	"context"
	"errors"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository/repoerrs"
)

type Portfolio struct {
	repos *repository.Repositories
}

func (p *Portfolio) AddStock(ctx context.Context, userId, companyId int) error {
	err := p.repos.Portfolio.Create(ctx, userId, companyId, 1)

	if err != nil {
		if errors.Is(err, repoerrs.ErrAlreadyExists) {
			return p.repos.Portfolio.AddStock(ctx, userId, companyId)
		}
	}

	return err
}

func (p *Portfolio) RemoveStock(ctx context.Context, userId, companyId int) error {
	return p.repos.Portfolio.RemoveStock(ctx, userId, companyId)
}

func NewPortfolio(repos *repository.Repositories) *Portfolio {
	return &Portfolio{repos: repos}
}
