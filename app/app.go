package app

import (
	"context"
	"github.com/modern-questions-team-13/orange-stock-market/config"
	"github.com/modern-questions-team-13/orange-stock-market/internal/controller"
	"github.com/modern-questions-team-13/orange-stock-market/internal/database"
	"github.com/modern-questions-team-13/orange-stock-market/internal/repository"
	"github.com/modern-questions-team-13/orange-stock-market/internal/service"
	"github.com/modern-questions-team-13/orange-stock-market/internal/tracer"
	"github.com/modern-questions-team-13/orange-stock-market/utility/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	oslog "log"
	"net"
	"net/http"
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

	// tracer
	tracerCfg := tracer.GetBaseConfig(cfg.Name)

	tr, closer, err := tracerCfg.NewTracer()
	if err != nil {
		log.Fatal().Err(err).Msg("tracer")
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tr)

	log.Info().Msg("initiating repositories...")
	repos := repository.NewRepositories(pg)

	log.Info().Msg("initiating services...")
	// service
	serv := service.NewServices(repos)

	//ttl serv
	serv.Ttl.Exec(context.Background())

	log.Info().Msg("initiating handler...")
	// handler
	h := controller.NewHandler(serv)

	// router
	log.Info().Msg("initiating router")

	router := controller.NewRouter(h)

	log.Info().Str("port", cfg.Port).Msg("starting listen server")
	oslog.Fatal(http.ListenAndServe(net.JoinHostPort("", cfg.Port), router.Router))
}
