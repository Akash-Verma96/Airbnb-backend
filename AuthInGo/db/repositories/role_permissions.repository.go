package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type RolePermissionRepository interface{
	GetRolePermissionById(id int64) (*models.RolePermission, error)
	GetRolePermissionByRoleId(roleId int64) ([]*models.Permission, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	GetAllRolePermission() ([]*models.RolePermission, error)
}

type RolePermissionRepositoryImpl struct{
	db *sql.DB
}

func NewRolePermissionRepository(_db *sql.DB) RolePermissionRepository{
	return &RolePermissionRepositoryImpl{
		db: _db,
	}
}

func (rpr *RolePermissionRepositoryImpl) GetRolePermissionById(id int64) (*models.RolePermission, error){

	query := "SELECT * FROM role_permissions where id = ?"

	row, err := rpr.db.Query(query, id);

	if err != nil {
		return nil, err
	}

	rolePermission := &models.RolePermission{}

	err = row.Scan(&rolePermission.Id,&rolePermission.RoleId,&rolePermission.PermissionId,&rolePermission.CreatedAt,&rolePermission.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No User found with this email!")
			return nil, err
		} else {
			fmt.Println("Error while fetching the user", err)
			return nil, err
		}
	}

	return rolePermission, nil
}

func (rpr *RolePermissionRepositoryImpl) GetRolePermissionByRoleId(roleId int64) ([]*models.Permission, error){
	query := `
		SELECT p.id, p.name, p.description, p.resource, p.action, p.created_at, p.updated_at
		FROM role_permissions rp
		INNER JOIN permissions p ON rp.permission_id = p.id
		WHERE rp.role_id = ?`

	rows, err := rpr.db.Query(query,roleId)

	if err != nil {
		fmt.Println("Error while fetching User Permission!", err)
		return nil,err
	}

	defer rows.Close()


	var permissions []*models.Permission

	for rows.Next() {
		permission := &models.Permission{}

		err := rows.Scan(&permission.Id,&permission.Name,&permission.Description,&permission.Resource,&permission.Action,&permission.CreatedAt,&permission.UpdatedAt)

		if err != nil {
			return nil,err
		}

		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil,err
	}

	return permissions,nil
}

func (rpr *RolePermissionRepositoryImpl) AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error){

	query := "INSERT INTO role_permissions (role_id, permission_id) VALUES (?,?)"

	rows, err := rpr.db.Exec(query,roleId,permissionId)

	if err != nil {
		fmt.Println("Error while Adding permissions to role", err)
		return nil, err
	}

	id, err := rows.LastInsertId()

	if err != nil {
		return nil,err
	}


	return &models.RolePermission{
		Id: id,
		RoleId: roleId,
		PermissionId: permissionId,
		CreatedAt: "",
		UpdatedAt: "",
	}, nil
}

func (rpr *RolePermissionRepositoryImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error{
	query := "DELETE FROM role_permissons where role_id = ? AND permission_id = ?"

	_, err := rpr.db.Exec(query,roleId,permissionId)

	if err != nil {
		fmt.Println("Error while removing role permission", err)
		return err
	}

	return nil
}

func (rpr *RolePermissionRepositoryImpl) GetAllRolePermission() ([]*models.RolePermission, error){
	query := "SELECT * FROM role_permissons"

	rows, err := rpr.db.Query(query)

	if err != nil {
		return nil,err
	}

	defer rows.Close()

	var rolePermissions []*models.RolePermission

	for rows.Next() {
		rolePermission := &models.RolePermission{}

		err := rows.Scan(&rolePermission.Id,&rolePermission.RoleId,&rolePermission.PermissionId,&rolePermission.CreatedAt,&rolePermission.UpdatedAt)

		if err != nil {
			return nil,err
		}

		rolePermissions = append(rolePermissions, rolePermission)
	}

	if err := rows.Err(); err != nil{
		return nil,err
	}

	return nil,nil
}