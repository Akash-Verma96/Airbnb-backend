package router

import (
	"AuthInGo/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Router interface {
	Register(r chi.Router)
}


func SetUpRouter(UserRouter Router, RoleRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger)

	chiRouter.Use(cors.Handler(cors.Options{
		
		AllowedOrigins:   []string{"https://*", "http://localhost:5173"},
		
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, 
  	}))


	chiRouter.HandleFunc("/HotelService/*", utils.ProxyToService("http://localhost:3002", "/HotelService"))
	chiRouter.HandleFunc("/BookingService/*", utils.ProxyToService("http://localhost:3001", "/BookingService"))

	UserRouter.Register(chiRouter)
	RoleRouter.Register(chiRouter)

	return chiRouter
}