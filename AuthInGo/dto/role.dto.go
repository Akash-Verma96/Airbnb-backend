package dto

import (

)

type CreateRoleDTO struct {
	Name string `json:"name" validate:"required,min=3"`
	Description string `json:"description" validate:"required,max=50"`
}

type UpdateRoleDTO struct {
    Name        string `json:"name" validate:"omitempty,min=3"`
    Description string `json:"description" validate:"omitempty,max=50"`
}

type AssignRoleDTO struct {
    UserId int64 `json:"userId" validate:"required"`
    RoleId int64 `json:"roleId" validate:"required"`
}

// AssignPermissionDTO, RemovePermissionDTO are used for assigning and removing permissions to/from roles
type AssignPermissionDTO struct {
    Id          int64  `json:"id" validate:"required"`
    PermissionId        int64 `json:"permissionId" validate:"required"`
}