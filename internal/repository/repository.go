package repository

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository/pgx"
)

type User interface {
	Create(ctx context.Context, login string, wealth int) (id int, err error)
	Withdraw(ctx context.Context, id int, wealth int) error
	TopUp(ctx context.Context, id int, wealth int) error
}

type Sale interface {
	Create(ctx context.Context, userId, companyId int, price int) (model.Sale, error)
	Get(ctx context.Context, id int) (model.Sale, error)
	Delete(ctx context.Context, id int) (model.Sale, error)
	GetSales(ctx context.Context, companyId int, maxPrice int, limit uint64) (id []int, err error)
	GetAllSales(ctx context.Context, companyId int, limit, offset uint64) (price []int, err error)
}

type Portfolio interface {
	AddStock(ctx context.Context, userId, companyId int) error
	RemoveStock(ctx context.Context, userId, companyId int) error
	Create(ctx context.Context, userId, companyId int, count int) error
}

type Buy interface {
	Create(ctx context.Context, userId, companyId int, price int) (model.Buy, error)
	Delete(ctx context.Context, id int) (model.Buy, error)
	Get(ctx context.Context, id int) (model.Buy, error)
	GetBuys(ctx context.Context, companyId int, minPrice int, limit uint64) (id []int, err error)
	GetAllBuys(ctx context.Context, companyId int, limit, offset uint64) (price []int, err error)
}

type Operation interface {
	Create(ctx context.Context, buyerId, sellerId, companyId int, price int) error
}

type Secret interface {
	SetToken(ctx context.Context, id int) (token string, err error)
	GetUserId(ctx context.Context, token string) (int, error)
}

type Repositories struct {
	User
	Sale
	Buy
	Operation
	Secret
	Portfolio
}

func NewRepositories(db *database.Postgres) *Repositories {
	return &Repositories{
		User:      pgx.NewUser(db),
		Sale:      pgx.NewSale(db),
		Buy:       pgx.NewBuy(db),
		Operation: pgx.NewOperation(db),
		Secret:    pgx.NewAuth(db),
		Portfolio: pgx.NewPortfolio(db),
	}
}
