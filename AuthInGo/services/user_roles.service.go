package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
)

type UserRoleService interface {
	AssignRoleToUserService(userId int64, roleId int64) error
	GetUserRolesService(userId int64) ([]*models.Role, error)
	RemoveRoleFromUserService(userId int64, roleId int64) error
}

type UserRoleServiceImpl struct {
	UserRoleRepository db.UserRoleRepository
}

func NewUserRoleService(_userRoleRepository db.UserRoleRepository) UserRoleService{
	return &UserRoleServiceImpl{
		UserRoleRepository: _userRoleRepository,
	}
}

func (urs *UserRoleServiceImpl) AssignRoleToUserService(userId int64, roleId int64) error {

	err := urs.UserRoleRepository.AssignRoleToUser(userId,roleId)

	return err
}

func (urs *UserRoleServiceImpl) GetUserRolesService(userId int64) ([]*models.Role, error){

	roles, err := urs.UserRoleRepository.GetUserRoles(userId)

	return roles,err
}

func (urs *UserRoleServiceImpl) RemoveRoleFromUserService(userId int64, roleId int64) error {

	err := urs.UserRoleRepository.RemoveRoleFromUser(userId,roleId)

	return err
}
