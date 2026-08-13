package app

import (
	config "AuthInGo/config/env"
	DBconfig "AuthInGo/config/db"
	"AuthInGo/controllers"
	repo "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
)

// config holds the configuration for server
type Config struct {
	Addr string
}

func NewConfig() Config {
	port := config.GetString("PORT",":8080")

	return Config{
		Addr: port,
	}
}
// Application holds the server details
type Application struct{
	Config Config
	Store repo.Storage
}

func NewApplication(config Config) *Application {
	return &Application{
		Config: config,
		Store: *repo.NewStorage(),
	}
}

// member function
func (app *Application) Run() error {

	db, err := DBconfig.SetUpDB()

	if err != nil {
		fmt.Println("Error while connecting DB")
	}

	ur := repo.NewUserRepository(db)
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	URouter := router.NewUserRouter(uc)

	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: router.SetUpRouter(URouter),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting Server in Go on port", app.Config.Addr)

	return server.ListenAndServe()
}
