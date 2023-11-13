package main

import (
	"context"
	"fmt"
	"github.com/modern-questions-team-13/orange-stock-market/bot/strategy"
	"time"
)

func main() {
	config := strategy.TwoBotStrategyConfig{
		Url: "http://localhost:9000/api",
		// id = 3
		FirstBot: strategy.BotConfig{
			Token:        "ba098c25f5a8321e9940447f551b5126f244e6ed",
			PriceForBuy:  strategy.NewRandomPrice(1, 100),
			PriceForSell: strategy.NewRandomPrice(1, 100),
			Interval:     time.Millisecond * 2,
		},
		// id = 4
		SecondBot: strategy.BotConfig{
			Token:        "fb7a973a860fee4183f0a23a4405ab3f2e098a76",
			PriceForBuy:  strategy.NewRandomPrice(1, 100),
			PriceForSell: strategy.NewRandomPrice(1, 100),
			Interval:     time.Millisecond * 2,
		},
		WorkTime: time.Second * 10,
	}

	bot := strategy.NewTwoBotStrategy(config)

	ctx, cancel := context.WithCancel(context.Background())
	bot.StartBot(cancel)

	<-ctx.Done()
	fmt.Println("Bot ustal i poshel spat'")
}
