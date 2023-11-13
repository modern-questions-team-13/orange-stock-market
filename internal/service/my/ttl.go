package my

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
	"github.com/rs/zerolog/log"
	"time"
)

type Ttl struct {
	repos *repository.Repositories
}

func (t *Ttl) Exec(ctx context.Context) {
	tickerBuy := time.NewTicker(time.Minute)

	tickerChanBuy := make(chan bool)

	go func() {
		for {
			select {
			case <-tickerChanBuy:
				return
			// interval task
			case tm := <-tickerBuy.C:
				log.Info().Str("time", tm.Format(time.ANSIC)).Msg("deleting expired buys")
				err := t.deleteExpiredBuys(ctx)

				if err != nil {
					log.Info().Err(err).Msg("delete expired buys")
					tickerBuy.Stop()
					tickerChanBuy <- true
				}
			}
		}
	}()

	tickerSale := time.NewTicker(time.Minute)
	tickerChanSale := make(chan bool)

	go func() {
		for {
			select {
			case <-tickerChanSale:
				return
			// interval task
			case tm := <-tickerSale.C:
				log.Info().Str("time", tm.Format(time.ANSIC)).Msg("deleting expired sales")
				err := t.deleteExpiredSales(ctx)

				if err != nil {
					log.Info().Err(err).Msg("delete expired sales")
					tickerBuy.Stop()
					tickerChanSale <- true
				}
			}
		}
	}()
}

func (t *Ttl) deleteExpiredBuys(ctx context.Context) error {
	bids, err := t.repos.Buy.DeleteExpired(ctx)

	if err != nil {
		return err
	}

	for _, bid := range bids {
		err = t.repos.User.TopUp(ctx, bid.UserId, bid.Price)
		if err != nil {
			return err
		}
	}

	return nil

}

func (t *Ttl) deleteExpiredSales(ctx context.Context) error {
	bids, err := t.repos.Sale.DeleteExpired(ctx)

	if err != nil {
		return err
	}

	for _, bid := range bids {
		err = t.repos.Portfolio.AddStock(ctx, bid.UserId, bid.CompanyId)
		if err != nil {
			return err
		}
	}

	return nil

}

func NewTtl(repos *repository.Repositories) *Ttl {
	return &Ttl{repos: repos}
}
