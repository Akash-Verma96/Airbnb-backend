package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	UserController *controllers.UserController
}


func NewUserRouter(_userController *controllers.UserController) Router{
	return &UserRouter{
		UserController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router){
	r.With(middlewares.CreateRequestValidator).Post("/signup", ur.UserController.CreateUser)
	r.With(middlewares.LoginRequestValidator, middlewares.RateLimiterMiddleware).Post("/login", ur.UserController.LoginUser)
	r.With(middlewares.JWTAuthMiddleware).Get("/profile", ur.UserController.GetUser)

	r.With(middlewares.JWTAuthMiddleware).Get("/all", ur.UserController.GetAllUser)
	r.With(middlewares.JWTAuthMiddleware).Delete("/", ur.UserController.DeleteUserById)
}