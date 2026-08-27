package dto

import (

)

type CreateRoleDTO struct {
	Name string `json:"name" validate:"required,min=3"`
	Description string `json:"description" validate:"required,max=50"`
}

type UpdateRoleDTO struct {
    Id          int64  `json:"id" validate:"required"`
    Name        string `json:"name" validate:"omitempty,min=3"`
    Description string `json:"description" validate:"omitempty,max=50"`
}
