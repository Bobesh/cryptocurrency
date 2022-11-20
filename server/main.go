package main

import (
	api2 "server/api"
	app2 "server/app"
	"server/config"
)

func main() {
	cfg := config.NewConfig()
	app := app2.NewApp(cfg.GetApiPath(), nil)
	api := api2.NewEchoApi(app, cfg.GetPort())
	api.Serve()
}
