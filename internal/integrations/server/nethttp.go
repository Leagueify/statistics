package server

import (
	"fmt"
	"net/http"

	"github.com/leagueify/statistics/handler"
	"github.com/leagueify/statistics/internal/config"
	"github.com/leagueify/statistics/internal/middleware"
)

type server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) Server {
	return &server{
		cfg: cfg,
	}
}

func (s *server) Start() {
	router := http.NewServeMux()
	statisticsRouter := handler.StatisticsRouter()

	router.Handle(
		"/statistics/", http.StripPrefix(
			"/statistics", statisticsRouter,
		),
	)

	// define server config
	statisticsServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Server.Port),
		Handler:      middleware.Logging(router),
		IdleTimeout:  s.cfg.Server.IdleTimeout,
		ReadTimeout:  s.cfg.Server.ReadTimeout,
		WriteTimeout: s.cfg.Server.WriteTimeout,
	}

	// show server banner
	showStartBanner()

	// start server
	if err := statisticsServer.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
