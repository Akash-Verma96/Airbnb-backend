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
	UpdateRole(id int64,payload dto.UpdateRoleDTO) (*models.Role, error)
	GetRolePermissions(roleId int64) ([]*models.Permission, error)
	AssignPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	AssignRoleToUser(userId int64, roleId int64) error
}

type RoleServiceImpl struct {
	roleRepository db.RoleRepository
	rolePermissionRepository db.RolePermissionRepository
	roleUserRepository db.UserRoleRepository
}

func NewRoleService(_roleRepository db.RoleRepository, _rolePermissionRepository db.RolePermissionRepository, _roleUserRepository db.UserRoleRepository) RoleService {
	return &RoleServiceImpl{
		roleRepository: _roleRepository,
		rolePermissionRepository: _rolePermissionRepository,
		roleUserRepository: _roleUserRepository,
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

func (s *RoleServiceImpl) UpdateRole(id int64, payload dto.UpdateRoleDTO) (*models.Role, error) {
	role, err := s.roleRepository.UpdateRole(id, payload.Name, payload.Description)

	return role, err
}

func (s *RoleServiceImpl) GetRolePermissions(roleId int64) ([]*models.Permission, error) {

	permissions, err := s.rolePermissionRepository.GetRolePermissionByRoleId(roleId)
	return permissions, err
}

func (s *RoleServiceImpl) AssignPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	return s.rolePermissionRepository.AddPermissionToRole(roleId, permissionId)
}

func (s *RoleServiceImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {
	return s.rolePermissionRepository.RemovePermissionFromRole(roleId, permissionId)
}

func (s *RoleServiceImpl) AssignRoleToUser(userId int64, roleId int64) error {
	err := s.roleUserRepository.AssignRoleToUser(userId, roleId)
	return err
}