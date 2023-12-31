package service

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/modern-questions-team-13/orange-stock-market/config"
	"github.com/modern-questions-team-13/orange-stock-market/internal/infrastructure/kafka"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
	"github.com/modern-questions-team-13/orange-stock-market/internal/service/my"
)

type User interface {
	Create(ctx context.Context, login string, wealth int) (token string, err error)
	Get(ctx context.Context, id int) (model.Info, error)
}

type Sale interface {
	Create(ctx context.Context, userId, companyId int, price int) error
	GetAllSalesByCompanyId(ctx context.Context, companyId int, limit, offset uint64) (price []int, err error)
}

type Buy interface {
	Create(ctx context.Context, userId, companyId int, price int) error
	GetAllBuysByCompanyId(ctx context.Context, companyId int, limit, offset uint64) (price []int, err error)
}

type Portfolio interface {
	AddStock(ctx context.Context, userId, companyId int) error
	RemoveStock(ctx context.Context, userId, companyId int) error
}

type Auth interface {
	GetUserId(ctx context.Context, token string) (int, bool)
}

type Company interface {
	GetAll(ctx context.Context) ([]model.Company, error)
}

type Ttl interface {
	Exec(ctx context.Context)
}

type KafkaSenderService interface {
	SendMessage(message kafka.RequestMessage) error
}

type Services struct {
	User
	Sale
	Buy
	Auth
	Portfolio
	Company
	Ttl
	KafkaSenderService
}

func NewServices(repos *repository.Repositories, producer sarama.SyncProducer, cfg *config.Config) *Services {
	return &Services{
		User:               my.NewUser(repos),
		Auth:               my.NewAuth(repos),
		Sale:               my.NewSale(repos),
		Buy:                my.NewBuy(repos),
		Portfolio:          my.NewPortfolio(repos),
		Company:            my.NewCompany(repos),
		Ttl:                my.NewTtl(repos),
		KafkaSenderService: my.NewKafkaSender(producer, cfg.RequestTopic),
	}
}
