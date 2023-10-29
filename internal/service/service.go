package service

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
)

type User interface {
	Create(ctx context.Context, login string, wealth int) (model.User, error) // TODO return token for auth
}

type Company interface {
	Create(ctx context.Context, name string) (model.Company, error)
}

type Sale interface {
	Create(ctx context.Context, userId, companyId int, price int) (model.Sale, error)
	Delete(ctx context.Context, id int) error
}

type Buy interface {
	Create(ctx context.Context, userId, companyId int, price int) (model.Buy, error)
	Delete(ctx context.Context, id int) error
}

type Operation interface {
	Create(ctx context.Context, userId, companyId int, price int) (model.Operation, error)
}

type Services struct {
	User
	Company
	Sale
	Buy
	Operation
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{}
}
