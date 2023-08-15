package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"temperature-sensor-receiver/internal/config"
	"temperature-sensor-receiver/internal/server"
)

func main() {
	var environment string

	flag.StringVar(&environment, "env", "prod", "Specify which environment to run in ('local' or 'prod')")
	flag.Parse()

	if environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	config.LoadConfig(environment)

	server, db := server.NewHTTPServer()
	defer db.Db.Close()

	server.Run(fmt.Sprintf("%s:%d", config.Config.Server.Host, config.Config.Server.Port))
}
