package strategy

import (
	"context"
	"fmt"
	"github.com/modern-questions-team-13/orange-stock-market/bot/handlers"
	"io"
	"math/rand"
	"sync"
	"time"
)

type Price interface {
	GetPrice() int64
}

type StaticPrice struct {
	price int64
}

func NewStaticPrice(price int64) *StaticPrice {
	return &StaticPrice{price: price}
}

func (p *StaticPrice) GetPrice() int64 {
	return p.price
}

type RandomPrice struct {
	priceMin int64
	priceMax int64
}

func NewRandomPrice(minPrice, maxPrice int64) *RandomPrice {
	return &RandomPrice{
		priceMin: minPrice,
		priceMax: maxPrice,
	}
}

func (p *RandomPrice) GetPrice() int64 {
	return rand.Int63n(p.priceMax-p.priceMin) + p.priceMin
}

type Bot struct {
	hand         handlers.BotHandler
	priceForBuy  Price
	priceForSell Price
	interval     time.Duration
}

type BotConfig struct {
	Token        string
	PriceForBuy  Price
	PriceForSell Price
	Interval     time.Duration
}

func NewBot(url string, cfg BotConfig) *Bot {
	return &Bot{
		hand:         handlers.NewBotHandler(url, cfg.Token),
		priceForBuy:  cfg.PriceForBuy,
		priceForSell: cfg.PriceForSell,
		interval:     cfg.Interval,
	}
}

func (s *Bot) buyAllOnce(waiter chan<- interface{}) {
	defer close(waiter)
	companies, err := s.hand.GetCompanies()
	if err != nil {
		fmt.Println("Problem with connecting", err.Error())
		return
	}

	wg := sync.WaitGroup{}
	for _, comp := range companies {
		wg.Add(1)
		go func(company handlers.Symbol) {

			count := 10
			for count > 0 {
				code, body, err := s.hand.LimitPriceBuy(company.Id, s.priceForBuy.GetPrice())
				if code == 429 {
					time.Sleep(s.interval)
				} else if err != nil {
					fmt.Println(err.Error(), "Company", company.Id, "Code Error", code)
				} else {

					res, _ := io.ReadAll(body)
					fmt.Println(code, string(res), company.Ticker, company.Id)
					count--
				}
			}
			wg.Done()
		}(comp)
	}
	wg.Wait()
	fmt.Println("all company start to buy")
}

func (s *Bot) sellAllOnce(waiter chan<- interface{}) {
	defer close(waiter)
	info, err := s.hand.GetInfo()
	if err != nil {
		fmt.Println("Problem with connecting", err.Error())
		return
	}
	wg := sync.WaitGroup{}
	for _, val := range info.Assets {
		wg.Add(1)
		go func(ass handlers.Asset) {
			for i := int64(0); i < ass.Quantity; i++ {
				for true {
					code, body, err := s.hand.LimitPriceSell(ass.Id, s.priceForSell.GetPrice())
					if code != 200 || err != nil {
						if err != nil {
							res, _ := io.ReadAll(body)
							fmt.Println("Code", code, "Error", err, "Msg", res)
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

type TwoBotStrategy struct {
	firstBot  *Bot
	secondBot *Bot
	workTime  time.Duration
}

type TwoBotStrategyConfig struct {
	Url       string
	FirstBot  BotConfig
	SecondBot BotConfig
	WorkTime  time.Duration
}

func (t *TwoBotStrategy) StartBot(cancelFunc context.CancelFunc) {
	timer := time.NewTimer(t.workTime)

	for {
		select {
		case <-timer.C:
			cancelFunc()
			return
		default:
			bot1 := make(chan interface{})
			bot2 := make(chan interface{})
			go t.firstBot.buyAllOnce(bot1)
			go t.secondBot.sellAllOnce(bot2)
			<-bot1
			<-bot2
		}
	}
}

func NewTwoBotStrategy(cfg TwoBotStrategyConfig) *TwoBotStrategy {
	return &TwoBotStrategy{
		firstBot:  NewBot(cfg.Url, cfg.FirstBot),
		secondBot: NewBot(cfg.Url, cfg.SecondBot),
		workTime:  cfg.WorkTime,
	}
}
