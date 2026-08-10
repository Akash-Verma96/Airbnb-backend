package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
)


func main(){
	config.Load() // env loaded


	cfg := app.NewConfig() // port assinged
	app := app.NewApplication(cfg) // server config done
	app.Run() // server up
	
}
