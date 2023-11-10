package my

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
)

type Sale struct {
	repos *repository.Repositories
}

func (s *Sale) Create(ctx context.Context, userId, companyId int, price int) (model.Sale, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Sale) Delete(ctx context.Context, id int) error {
	return s.repos.Sale.Delete(ctx, id)
}

func (s *Sale) GetAllSales(ctx context.Context, companyId int, limit, offset uint64) (price []int, err error) {
	return s.repos.Sale.GetAllSales(ctx, companyId, limit, offset)
}

func NewSale(repos *repository.Repositories) *Sale {
	return &Sale{repos: repos}
}
