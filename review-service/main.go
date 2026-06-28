package main

import (
	"Review/app"
	config "Review/config/env"
)

func main() {

	config.Load()

	cfg := app.NewConfig() // Set the server to listen on port 8081
	app := app.NewApplication(cfg)

	app.Run()
}