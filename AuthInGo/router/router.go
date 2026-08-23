package router

import (
	"AuthInGo/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface {
	Register(r chi.Router)
}


func SetUpRouter(UserRouter Router, RoleRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Use(middleware.Logger)
	// { targetBaseUrl , pathPrefix }

	// chiRouter.HandleFunc("/fakestoreservice/*", utils.ProxyToService("https://fakestoreapi.com", "/fakestoreservice"))
	chiRouter.HandleFunc("/HotelService/*", utils.ProxyToService("http://localhost:3002", "/HotelService"))
	chiRouter.HandleFunc("/BookingService/*", utils.ProxyToService("http://localhost:3001", "/BookingService"))

	UserRouter.Register(chiRouter)
	RoleRouter.Register(chiRouter)

	return chiRouter
}