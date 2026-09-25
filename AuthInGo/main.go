package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
	scheduler "AuthInGo/scheduler"
)


func main(){
	config.Load() // env loaded


	cfg := app.NewConfig() // port assinged
	app := app.NewApplication(cfg) // server config done
	scheduler.ApiScheduler()
	app.Run() // server up
	
}
