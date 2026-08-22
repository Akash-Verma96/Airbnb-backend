package controllers

import (
	"AuthInGo/services"
	"AuthInGo/utils"
	"net/http"
)

type UserRoleController struct {
	UserRoleService services.UserRoleService
}

func NewUserRoleController(_userRoleService services.UserRoleService) *UserRoleController{
	return &UserRoleController{
		UserRoleService: _userRoleService,
	}
}

func (urc *UserRoleController) AssignRoleToUser(w http.ResponseWriter, r *http.Request){

	err := urc.UserRoleService.AssignRoleToUserService(3,2)


	utils.WriteJsonSuccessResponse(w,http.StatusOK,"User Role Assinged Successfully!", err)
}

func (urc *UserRoleController) GetUserRoles(w http.ResponseWriter, r *http.Request){

	roles, err := urc.UserRoleService.GetUserRolesService(3)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error while fetching the User Role!", err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK,"User Roles Fetched Successfully!", roles)
}

func (urc *UserRoleController) RemoveRoleFromUser(w http.ResponseWriter, r *http.Request){

	err := urc.UserRoleService.RemoveRoleFromUserService(3,2)

	utils.WriteJsonSuccessResponse(w,http.StatusOK, "User Role removed Successfully!", err)
}