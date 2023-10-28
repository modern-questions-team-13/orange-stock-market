package repository

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository/pgx"
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
	Create(ctx context.Context, userId, companyId int, price int, opType model.OperationType) (model.Operation, error)
}

type Secret interface {
	// TODO
}

type Repositories struct {
	User
	Company
	Sale
	Buy
	Operation
	//TODO Secret
}

func NewRepositories(db *database.Postgres) *Repositories {
	return &Repositories{
		User:      pgx.NewUser(db),
		Company:   pgx.NewCompany(db),
		Sale:      pgx.NewSale(db),
		Buy:       pgx.NewBuy(db),
		Operation: pgx.NewOperation(db),
	}
}
