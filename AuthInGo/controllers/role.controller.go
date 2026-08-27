package controllers

import (
	"AuthInGo/dto"
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


func (rc *RoleController) GetRoleByName(w http.ResponseWriter, r *http.Request) {

	var payload struct {
		Name string `json:"name"`
	}

	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Bad Input", jsonErr)
		return
	}

	role, err := rc.roleService.GetRoleByName(payload.Name)

	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Error while Fetching Role", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role Fetched Successfully!", role)
}

func (rc *RoleController) CreateRole(w http.ResponseWriter, r *http.Request){
	payload := r.Context().Value("payload-key").(dto.CreateRoleDTO)

	role, err := rc.roleService.CreateRole(payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Error while creating role", err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Role created Successfully!",role)
}

func (rc *RoleController) DeleteById(w http.ResponseWriter,r *http.Request){

	var payload struct {
		Id int64 `json:"id"`
	}

	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Bad Input", jsonErr)
		return
	}

	err := rc.roleService.DeleteRoleById(payload.Id)

	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Error while Deleting Role", err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Role Deleted Successfully!", "nil")
}



func (rc *RoleController) UpdateRole(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value("payload-key").(dto.UpdateRoleDTO)

	role, err := rc.roleService.UpdateRole(payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Error While updating role", err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Role updated Successfully!",role)
}

func (rc *RoleController) GetRolePermissions(w http.ResponseWriter, r *http.Request){

	var payload struct {
		Id int64 `json:"id"`
	}


	if jsonErr := utils.ReadJsonBody(r,&payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Bad Input", jsonErr)
		return
	}

	permissions, err := rc.roleService.GetRolePermissions(payload.Id)

	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Error while getting permissions!", err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK,"All the Permissions Found!",permissions)
}

func (rc *RoleController) AddPermissionToRole(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Id int64 `json:"id"`
		PermissionId int64 `json:"permissionId"`
	}

	if jsonErr := utils.ReadJsonBody(r,&payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Bad Input", jsonErr)
		return
	}

	permission, err := rc.roleService.AddPermissionToRole(payload.Id,payload.PermissionId)

	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Error while Adding Permission!", err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Permission Added Successfull!", permission)
}