package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
)

type RoleService interface{
	GetRoleById(id int64) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(name string, description string) (*models.Role, error)
	DeleteRoleById(id int64) error
	UpdateRole(id int64, name string, description string) (*models.Role, error)
	GetRolePermissions(roleId int64) ([]*models.RolePermission, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
}

type RoleServiceImpl struct {
	roleRepository db.RoleRepository
	rolePermissionRepository db.RolePermissionRepository
}

func NewRoleService(roleRepository db.RoleRepository) RoleService {
	return &RoleServiceImpl{
		roleRepository: roleRepository,
	}
}


func (s *RoleServiceImpl) GetRoleById (id int64) (*models.Role, error) {

	role, err := s.roleRepository.GetRoleById(id)

	return role, err
}

func (s *RoleServiceImpl) GetRoleByName(name string) (*models.Role, error) {
	role, err := s.roleRepository.GetRoleByName(name)

	return role, err
}

func (s *RoleServiceImpl) GetAllRoles() ([]*models.Role, error) {
	roles, err := s.roleRepository.GetAllRoles()

	return roles, err
}

func (s *RoleServiceImpl) CreateRole(name string, description string) (*models.Role, error) {
	role, err := s.roleRepository.CreateRole(name, description)

	return role, err
}


func (s *RoleServiceImpl) DeleteRoleById(id int64) error {
	err := s.roleRepository.DeleteRoleById(id)

	return err
}

func (s *RoleServiceImpl) UpdateRole(id int64, name string, description string) (*models.Role, error) {
	role, err := s.roleRepository.UpdateRole(id, name, description)

	return role, err
}

func (s *RoleServiceImpl) GetRolePermissions(roleId int64) ([]*models.RolePermission, error) {

	permissions, err := s.rolePermissionRepository.GetRolePermissionByRoleId(roleId)
	return permissions, err
}

func (s *RoleServiceImpl) AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	return s.rolePermissionRepository.AddPermissionToRole(roleId, permissionId)
}