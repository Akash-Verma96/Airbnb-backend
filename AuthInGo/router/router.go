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
		
		AllowedOrigins:   []string{"*"},
		
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, 
  	}))


	chiRouter.HandleFunc("/HotelService/*", utils.ProxyToService("https://airbnb-hotel-0job.onrender.com", "/HotelService"))
	chiRouter.HandleFunc("/BookingService/*", utils.ProxyToService("https://airbnb-backend-glo7.onrender.com", "/BookingService"))

	// chiRouter.HandleFunc("/HotelService/*", utils.ProxyToService("http://localhost:3001", "/HotelService"))
	// chiRouter.HandleFunc("/BookingService/*", utils.ProxyToService("http://localhost:3002", "/BookingService"))

	UserRouter.Register(chiRouter)
	RoleRouter.Register(chiRouter)

	return chiRouter
}