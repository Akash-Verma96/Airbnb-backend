package router

import (
	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func SetUpRouter(UserRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()


	UserRouter.Register(chiRouter)

	return chiRouter
}