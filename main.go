package main

import (
	"invest/app"
	"invest/app/config"
	"invest/app/services"
)

func main() {
	Config := config.Loaders{}

	Config.LoadEnv()
	services.LoadDB()
	app.Run()
}
