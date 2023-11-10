package my

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
)

type Buy struct {
	repos *repository.Repositories
}

func (b *Buy) Create(ctx context.Context, userId, companyId int, price int) (model.Buy, error) {
	//TODO implement me
	panic("implement me")
}

func (b *Buy) Delete(ctx context.Context, id int) error {
	return b.repos.Buy.Delete(ctx, id)
}

func (b *Buy) GetAllBuys(ctx context.Context, companyId int, limit, offset uint64) (price []int, err error) {
	return b.repos.GetAllBuys(ctx, companyId, limit, offset)
}

func NewBuy(repos *repository.Repositories) *Buy {
	return &Buy{repos: repos}
}
