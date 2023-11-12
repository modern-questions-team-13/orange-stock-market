package strategy

import (
	"context"
	"fmt"
	"github.com/modern-questions-team-13/orange-stock-market/bot/handlers"
	"sync"
	"time"
)

type Strategy interface {
	StartBot(cancelFunc context.CancelFunc)
}

type SimpleStrategy struct {
	hand         handlers.BotHandler
	priceForSell int64
	priceForBuy  int64
	interval     time.Duration
	workTime     time.Duration
}

type SimpleStrategyConfig struct {
	Url          string
	Token        string
	PriceForBuy  int64
	PriceForSell int64
	Interval     time.Duration
	WorkTime     time.Duration
}

func NewSimpleStrategy(cfg SimpleStrategyConfig) *SimpleStrategy {
	return &SimpleStrategy{
		hand:         handlers.NewBotHandler(cfg.Url, cfg.Token),
		priceForSell: cfg.PriceForSell,
		priceForBuy:  cfg.PriceForBuy,
		interval:     cfg.Interval,
		workTime:     cfg.WorkTime,
	}
}

func (s *SimpleStrategy) StartBot(cancelFunc context.CancelFunc) {
	timer := time.NewTimer(s.workTime)

	for {
		select {
		case <-timer.C:
			cancelFunc()
			return
		default:
			s.buyAllOnce()
			s.sellAllOnce()
		}
	}
}

func (s *SimpleStrategy) buyAllOnce() {
	companies, err := s.hand.GetCompanies()
	for err != nil {
		time.Sleep(s.interval)
		companies, err = s.hand.GetCompanies()
	}

	wg := sync.WaitGroup{}
	for _, comp := range companies {
		wg.Add(1)
		go func(company handlers.Symbol) {
			for true {
				code, err := s.hand.LimitPriceBuy(company.Id, s.priceForBuy)
				if code == 429 {
					time.Sleep(s.interval)
				} else if err != nil {
					fmt.Println(err.Error(), "Company", company.Id, "Code Error", code)
				} else {
					fmt.Println(code, company.Ticker, company.Id)
					break
				}
			}
			wg.Done()
		}(comp)
	}
	wg.Wait()
	fmt.Println("all company start to buy")
}

func (s *SimpleStrategy) sellAllOnce() {
	info, err := s.hand.GetInfo()
	for err != nil {
		time.Sleep(s.interval)
		info, err = s.hand.GetInfo()
	}
	wg := sync.WaitGroup{}
	for _, val := range info.Assets {
		wg.Add(1)
		go func(ass handlers.Asset) {
			for i := int64(0); i < ass.Quantity; i++ {
				for true {
					code, err := s.hand.LimitPriceSell(ass.Id, s.priceForSell)
					if code != 200 || err != nil {
						if err != nil {
							fmt.Println("Code", code, "Error", err)
						}
						time.Sleep(s.interval)
					} else {
						break
					}
				}
			}
			fmt.Println(ass.Id, ass.Quantity)
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println("all company start to sell")
}
