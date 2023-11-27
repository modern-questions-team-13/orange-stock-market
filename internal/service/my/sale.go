package my

import (
	"context"
	"errors"
	"fmt"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository/repoerrs"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
)

type Sale struct {
	repos *repository.Repositories
}

func (s *Sale) Create(ctx context.Context, userId, companyId int, price int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "app: create sale")
	defer span.Finish()

	err := s.repos.Portfolio.RemoveStock(ctx, userId, companyId)

	if err != nil {
		return err
	}

	ok, err := s.trySellImmediately(ctx, userId, companyId, price)

	if err != nil {
		if !ok {
			log.Err(err).Msg("immediate sell")
			return err
		}

		_, _err := s.repos.Sale.Create(ctx, userId, companyId, price)

		if _err != nil {
			return err
		}

		log.Info().Msg(err.Error())
	}

	return nil
}

func (s *Sale) GetAllSalesByCompanyId(ctx context.Context, companyId int, limit, offset uint64) (price []int, err error) {
	return s.repos.GetAllSales(ctx, companyId, limit, offset)
}

func (s *Sale) GetAllSales(ctx context.Context, companyId int, limit, offset uint64) (price []int, err error) {
	return s.repos.Sale.GetAllSales(ctx, companyId, limit, offset)
}

func (s *Sale) trySellImmediately(ctx context.Context, userId int, companyId int, price int) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "app: try sell immediately")
	defer span.Finish()
	buys, err := s.repos.Buy.GetBuys(ctx, companyId, price, tryBuyLimit)

	if err != nil {
		return false, err
	}

	span, ctx = opentracing.StartSpanFromContext(ctx, "app: range in sell")
	for _, buyId := range buys {
		if buy, _err := s.repos.Buy.Delete(ctx, buyId); _err == nil {
			err = s.makeSellOperation(ctx, userId, buy, price)

			if err != nil {
				return false, err
			}

			return true, nil
		}
	}
	span.Finish()

	return true, fmt.Errorf("not finding matching sales")
}

func (s *Sale) makeSellOperation(ctx context.Context, sellerId int, buy model.Buy, price int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "app: makes sell operation")
	defer span.Finish()

	if buy.Price > price {
		err := s.repos.User.TopUp(ctx, buy.UserId, buy.Price-price)

		if err != nil {
			return err
		}
	}

	err := s.addStock(ctx, buy.UserId, buy.CompanyId)

	if err != nil {
		return err
	}

	err = s.repos.User.TopUp(ctx, sellerId, price)

	if err != nil {
		return err
	}

	return s.repos.Operation.Create(ctx, buy.UserId, sellerId, buy.CompanyId, price)
}

func (s *Sale) addStock(ctx context.Context, userId, companyId int) error {
	err := s.repos.Portfolio.Create(ctx, userId, companyId, 1)

	if err != nil {
		if errors.Is(err, repoerrs.ErrAlreadyExists) {
			return s.repos.Portfolio.AddStock(ctx, userId, companyId)
		}
	}

	return err
}

func NewSale(repos *repository.Repositories) *Sale {
	return &Sale{repos: repos}
}
