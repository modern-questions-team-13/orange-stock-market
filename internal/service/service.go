package service

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
	"github.com/modern-questions-team-13/orange-stock-market/internal/service/my"
)

type User interface {
	Create(ctx context.Context, login string, wealth int) (token string, err error)
}

type Sale interface {
	Create(ctx context.Context, userId, companyId int, price int) (model.Sale, error)
	Delete(ctx context.Context, id int) error
	GetAllSales(ctx context.Context, companyId int, limit, offset uint64) (price []int, err error)
}

type Buy interface {
	Create(ctx context.Context, userId, companyId int, price int) (model.Buy, error)
	Delete(ctx context.Context, id int) error
}

type Operation interface {
	Create(ctx context.Context, userId, companyId int, price int) (model.Operation, error)
}

type Auth interface {
	GetUserId(ctx context.Context, token string) (int, bool)
}

type Services struct {
	User
	Sale
	Buy
	Operation
	Auth
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		User: my.NewUser(repos),
		Auth: my.NewAuth(repos),
	}
}
