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
	
	r.Get("/roles/{id}", rr.roleController.GetRoleById)
	r.Get("/roles", rr.roleController.GetAllRoles)
	r.With(middlewares.CreateRoleRequestValidator).Post("/createRole", rr.roleController.CreateRole)
	r.With(middlewares.UpdateRoleRequestValidator).Patch("/roles/{id}", rr.roleController.UpdateRole)
	r.Get("/roleByName", rr.roleController.GetRoleByName)

	r.Get("/role/{id}/permissions", rr.roleController.GetRolePermissions)
	r.With(middlewares.AssignPermissionRequestValidator).Post("/assignpermissionToRole", rr.roleController.AssignPermissionToRole)
	r.With(middlewares.RemovePermissionRequestValidator).Delete("/removePermissionFromRole", rr.roleController.RemovePermissionFromRole)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAllRoles("admin")).Post("/roles/{userId}/assign/{roleId}", rr.roleController.AssignRoleToUser)
	
	r.Delete("/roles/{id}", rr.roleController.DeleteById)
} 