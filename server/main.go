package main

import (
	api2 "server/api"
	apps "server/apps"
	"server/config"
)

func main() {
	cfg := config.NewConfig()
	cryptoApp := apps.NewCryptoApp(cfg.GetApiPath(), nil)
	hashApp := apps.NewHashApp(cfg.GetFilesDir())
	api := api2.NewEchoApi(cryptoApp, hashApp, cfg.GetPort())
	api.Serve()
}
