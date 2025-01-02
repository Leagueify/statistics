package main

import (
	"github.com/leagueify/statistics/internal/config"
	"github.com/leagueify/statistics/internal/integrations/server"
)

func main() {
	cfg := config.GetConfig()

	server.NewServer(cfg).Start()
}
