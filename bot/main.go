package main

import (
	"context"
	"fmt"
	"github.com/modern-questions-team-13/orange-stock-market/bot/strategy"
	"time"
)

func main() {
	bot := strategy.NewSimpleStrategy(strategy.SimpleStrategyConfig{
		Url:          "http://localhost:9000/api",
		Token:        "fb7a973a860fee4183f0a23a4405ab3f2e098a76",
		PriceForBuy:  1000,
		PriceForSell: 2,
		Interval:     time.Millisecond * 2,
		WorkTime:     time.Second * 10,
	})
	ctx, cancel := context.WithCancel(context.Background())
	bot.StartBot(cancel)
	<-ctx.Done()
	fmt.Println("Bot ustal i poshel spat'")
}
