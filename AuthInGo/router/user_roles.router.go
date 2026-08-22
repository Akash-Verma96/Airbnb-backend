package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRoleRouter struct {
	UserRoleController *controllers.UserRoleController
}

func NewUserRoleRouter(_userRoleController *controllers.UserRoleController) Router{
	return &UserRoleRouter{
		UserRoleController: _userRoleController,
	}
}

func (urr *UserRoleRouter) Register(r chi.Router){
	r.Post("/assignRole", urr.UserRoleController.AssignRoleToUser)
	r.Get("/getRole", urr.UserRoleController.GetUserRoles)

	r.Delete("/removeRole", urr.UserRoleController.RemoveRoleFromUser)
}