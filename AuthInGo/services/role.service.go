package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
)

type RoleService interface{
	GetRoleById(id int64) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(payload dto.CreateRoleDTO) (*models.Role, error)
	DeleteRoleById(id int64) error
	UpdateRole(payload dto.UpdateRoleDTO) (*models.Role, error)
	GetRolePermissions(roleId int64) ([]*models.Permission, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
}

type RoleServiceImpl struct {
	roleRepository db.RoleRepository
	rolePermissionRepository db.RolePermissionRepository
}

func NewRoleService(_roleRepository db.RoleRepository, _rolePermissionRepository db.RolePermissionRepository) RoleService {
	return &RoleServiceImpl{
		roleRepository: _roleRepository,
		rolePermissionRepository: _rolePermissionRepository,
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

func (s *RoleServiceImpl) CreateRole(payload dto.CreateRoleDTO) (*models.Role, error) {
	role, err := s.roleRepository.CreateRole(payload.Name, payload.Description)

	return role, err
}


func (s *RoleServiceImpl) DeleteRoleById(id int64) error {
	err := s.roleRepository.DeleteRoleById(id)

	return err
}

func (s *RoleServiceImpl) UpdateRole(payload dto.UpdateRoleDTO) (*models.Role, error) {
	role, err := s.roleRepository.UpdateRole(payload.Id, payload.Name, payload.Description)

	return role, err
}

func (s *RoleServiceImpl) GetRolePermissions(roleId int64) ([]*models.Permission, error) {

	permissions, err := s.rolePermissionRepository.GetRolePermissionByRoleId(roleId)
	return permissions, err
}

func (s *RoleServiceImpl) AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	return s.rolePermissionRepository.AddPermissionToRole(roleId, permissionId)
}