package controllers

import (
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoleController struct {
	roleService services.RoleService
}

func NewRoleController(_roleService services.RoleService) *RoleController {
	return &RoleController{
		roleService: _roleService,
	}
}


func (rc *RoleController) GetRoleById(w http.ResponseWriter, r *http.Request) {
	
	roleId := chi.URLParam(r, "id") // Extract role ID from URL parameters
	if roleId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}
	id, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	role, err := rc.roleService.GetRoleById(id)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch role", err)
		return
	}


	utils.WriteJsonSuccessResponse(w,http.StatusOK, "Role Found", role)
}


func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	
	roles, err := rc.roleService.GetAllRoles()
	
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to fetch roles", err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK, "All Roles Found", roles)
}