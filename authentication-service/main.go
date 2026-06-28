package main

import (
	"Authingo/app"

	config "Authingo/config/env"
	
)
func main(){
	config.Load()
	cfg:=app.NewConfig()
	app:=app.NewApplication(cfg)
	
	app.Run()
}