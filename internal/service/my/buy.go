package my

import (
	"context"
	"errors"
	"fmt"
	"github.com/modern-questions-team-13/orange-stock-market/internal/model"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository/repoerrs"
	"github.com/rs/zerolog/log"
)

const tryBuyLimit = 2

type Buy struct {
	repos *repository.Repositories
}

func (b *Buy) Create(ctx context.Context, userId, companyId int, price int) error {
	err := b.repos.User.Withdraw(ctx, userId, price)

	if err != nil {
		return err
	}

	ok, err := b.tryBuyImmediately(ctx, userId, companyId, price)

	if err != nil {
		if !ok {
			log.Err(err).Msg("immediate buy")
			return err
		}

		_, _err := b.repos.Buy.Create(ctx, userId, companyId, price)

		if _err != nil {
			return err
		}

		log.Info().Msg(err.Error())
	}

	return nil
}

func (b *Buy) GetAllBuysByCompanyId(ctx context.Context, companyId int, limit, offset uint64) (price []int, err error) {
	return b.repos.GetAllBuys(ctx, companyId, limit, offset)
}

func (b *Buy) tryBuyImmediately(ctx context.Context, userId, companyId int, price int) (bool, error) {
	sales, err := b.repos.Sale.GetSales(ctx, companyId, price, tryBuyLimit)

	if err != nil {
		return false, err
	}

	for _, saleId := range sales {
		if sale, _err := b.repos.Sale.Delete(ctx, saleId); _err == nil {
			err = b.makeBuyOperation(ctx, userId, sale, price)

			if err != nil {
				return false, err
			}

			return true, nil
		}
	}

	return true, fmt.Errorf("not finding matching sales")
}

func (b *Buy) addStock(ctx context.Context, userId, companyId int) error {
	err := b.repos.Portfolio.Create(ctx, userId, companyId, 1)

	if err != nil {
		if errors.Is(err, repoerrs.ErrAlreadyExists) {
			return b.repos.Portfolio.AddStock(ctx, userId, companyId)
		}
	}

	return err
}

func (b *Buy) makeBuyOperation(ctx context.Context, buyerId int, sale model.Sale, reservedMoney int) error {
	err := b.repos.User.TopUp(ctx, sale.UserId, sale.Price)

	if err != nil {
		_ = b.addStock(ctx, sale.UserId, sale.CompanyId)
		return err
	}

	if sale.Price < reservedMoney {
		err = b.repos.User.TopUp(ctx, buyerId, reservedMoney-sale.Price)

		if err != nil {
			return err
		}
	}

	err = b.addStock(ctx, buyerId, sale.CompanyId)

	if err != nil {
		return err
	}

	return b.repos.Operation.Create(ctx, buyerId, sale.UserId, sale.CompanyId, sale.Price)
}

func NewBuy(repos *repository.Repositories) *Buy {
	return &Buy{repos: repos}
}
