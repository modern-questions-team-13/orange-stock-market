package app

import (
	"context"
	"fmt"
	"github.com/modern-questions-team-13/orange-stock-market/config"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
	"github.com/modern-questions-team-13/orange-stock-market/internal/service"
	"github.com/modern-questions-team-13/orange-stock-market/utility/logger"
	"github.com/rs/zerolog/log"
	"sync"
)

func Run() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal().Err(err).Msg("config")
	}

	// logger
	logger.SetLevel(cfg.Level)

	// postgres
	log.Info().Msg("connecting to postgres...")
	pg, err := database.NewPostgres(cfg.Url)

	if err != nil {
		log.Fatal().Err(err).Msg("postgres connect")
	}

	defer pg.Close()

	log.Info().Msg("initiating repositories...")
	repos := repository.NewRepositories(pg)

	log.Info().Msg("initiating services...")
	// service
	serv := service.NewServices(repos)

	var wg = sync.WaitGroup{}

	var n = 100

	wg.Add(n)

	for i := 1; i <= n; i++ {
		go func(i int) {
			defer wg.Done()

			if i%2 == 0 {
				_err := serv.Buy.Create(context.Background(), 2, 1, 300)

				if _err != nil {
					fmt.Printf("buy: %s\n", _err.Error())
				}
			} else {
				_err := serv.Sale.Create(context.Background(), 1, 1, 10)

				if _err != nil {
					fmt.Printf("buy: %s\n", _err.Error())
				}
			}
		}(i)
	}

	wg.Wait()
	/*
		log.Info().Msg("initiating handler...")
		// handler
		h := controller.NewHandler(serv)

		// router
		log.Info().Msg("initiating router")

		router := controller.NewRouter(h)

		log.Info().Str("port", cfg.Port).Msg("starting listen server")
		oslog.Fatal(http.ListenAndServe(net.JoinHostPort("", cfg.Port), router.Router))
	*/
}
