package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	roleController *controllers.RoleController
}

func NewRoleRouter(_roleController *controllers.RoleController) Router{
	return &RoleRouter{
		roleController: _roleController,
	}
}

func (rr *RoleRouter) Register(r chi.Router){
	r.With(middlewares.CreateRoleRequestValidator).Post("/createRole", rr.roleController.CreateRole)

	r.Get("/roles/{id}", rr.roleController.GetRoleById)
	r.Get("/roles", rr.roleController.GetAllRoles)
	r.Get("/roleByName", rr.roleController.GetRoleByName)

	r.Delete("/", rr.roleController.DeleteById)
	r.With(middlewares.UpdateRoleRequestValidator).Patch("/", rr.roleController.UpdateRole)
	r.Get("/rolePermission", rr.roleController.GetRolePermissions)
	r.Post("/addPermission", rr.roleController.AddPermissionToRole)
} 