package main

import (
	"context"
	"fmt"
	"github.com/modern-questions-team-13/orange-stock-market/bot/strategy"
)

func main() {
	bot := strategy.NewSimpleStrategy(strategy.SimpleStrategyConfig{
		Url:          "",
		Token:        "",
		PriceForBuy:  0,
		PriceForSell: 0,
		Interval:     0,
		WorkTime:     0,
	})
	ctx, cancel := context.WithCancel(context.Background())
	bot.StartBot(cancel)
	<-ctx.Done()
	fmt.Println("Bot ustal i poshel spat'")
}
