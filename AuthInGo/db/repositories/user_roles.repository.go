package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
	"strings"
)

type UserRoleRepository interface {
	AssignRoleToUser(userId int64, roleId int64) error
	GetUserRoles(userId int64) ([]*models.Role, error)
	RemoveRoleFromUser(userId int64, roleId int64) error
	GetUserPermissions(userId int64) ([]*models.Permission, error) // to check the permissions of a user can be done by joining Role_Pemission table
	HasPermission(userId int64, permissionName string) (bool, error) //to check weather the user has permission to do set of task or not
	HasRole(userId int64, roleName string) (bool, error) // hasRole()
	HasRoles(userId int64, roleNames []string) (bool, error)
}

type UserRoleRepositoryImpl struct{
	db  *sql.DB
}

func NewUserRoleRepository(_db *sql.DB) UserRoleRepository{
	return &UserRoleRepositoryImpl{
		db: _db,
	}
}

func(ur *UserRoleRepositoryImpl) AssignRoleToUser(userId int64, roleId int64) error {

	query := "INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)"

	_, err := ur.db.Exec(query,userId,roleId)

	if err != nil {
		fmt.Println("Error While Executing the query", err)
		return err
	}

	return nil
}

func (ur *UserRoleRepositoryImpl) GetUserRoles(userId int64) ([]*models.Role, error) {

	query := `SELECT r.id, r.name, r.description, r.created_at, r.updated_at
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ?`

	rows, err := ur.db.Query(query,userId)

	if err != nil {
		fmt.Println("Error while fetching user Role", err)
		return nil,err
	}

	defer rows.Close()

	var roles []*models.Role


	for rows.Next() {
		var role = &models.Role{}
		//  Scan the database columns into the fields of the struct
		err := rows.Scan(&role.Id,&role.Name,&role.Description,&role.CreatedAt,&role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		// Append each role to your slice
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil,err
	}


	return roles,nil
}



func(ur *UserRoleRepositoryImpl) RemoveRoleFromUser(userId int64, roleId int64) error {

	query := "DELETE FROM user_roles WHERE user_id = ? AND role_id = ?"

	_, err := ur.db.Exec(query,userId,roleId)

	if err != nil {
		fmt.Println("Error while removing role", err)
		return err
	}

	return nil
}

func(ur *UserRoleRepositoryImpl) GetUserPermissions(userId int64) ([]*models.Permission, error) {
		query := `
		SELECT p.id, p.name, p.description, p.resource, p.action, p.created_at, p.updated_at
		FROM user_roles ur
		INNER JOIN role_permissions rp ON ur.role_id = rp.role_id
		INNER JOIN permissions p ON rp.permission_id = p.id
		WHERE ur.user_id = ?`

		rows, err := ur.db.Query(query,userId);

		if err != nil {
			fmt.Println("Error while fetching the user Permission!", err)
			return nil,err
		}

		defer rows.Close()

		var permissions []*models.Permission

		for rows.Next() {
			var permission = &models.Permission{}

			err := rows.Scan(&permission.Id,&permission.Name,&permission.Resource,&permission.Action,&permission.CreatedAt,&permission.UpdatedAt)

			if err != nil {
				return nil, err
			}

			permissions = append(permissions,permission)
		}


		if err := rows.Err(); err != nil {
			return nil, err
		}

	return permissions,nil
}

func (ur *UserRoleRepositoryImpl) HasPermission(userId int64, permissionName string) (bool, error) {
	query := `
		SELECT COUNT(*) > 0
		FROM user_roles ur
		INNER JOIN role_permissions rp ON ur.role_id = rp.role_id
		INNER JOIN permissions p ON rp.permission_id = p.id
		WHERE ur.user_id = ? AND p.name = ?`

	var exists bool
	err := ur.db.QueryRow(query, userId, permissionName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (ur *UserRoleRepositoryImpl) HasRole(userId int64, roleName string) (bool, error) {
	query := `
			SELECT COUNT(*) > 0
			FROM user_roles ur
			INNER JOIN roles r ON ur.role_id = r.id
			WHERE ur.user_id = ? AND r.name = ?`
	var exists bool
	err := ur.db.QueryRow(query, userId, roleName).Scan(&exists)
	if err != nil {
		return false, err
	}	
	return exists, nil
}


func (ur *UserRoleRepositoryImpl) HasRoles(userId int64, roleNames []string) (bool, error) {

	if len(roleNames) == 0 {
		return true, nil // If no roles are specified, return true
	}


	query := `
		SELECT COUNT(*) = ?
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ? AND r.name IN (?)
		GROUP BY ur.user_id`

	roleNamesStr := strings.Join(roleNames, ",")
	row := ur.db.QueryRow(query, len(roleNames), userId, roleNamesStr)

	var hasAllRoles bool
	if err := row.Scan(&hasAllRoles); err != nil {
		if err == sql.ErrNoRows {
			return false, nil // No roles found for the user
		}
		return false, err // Return any other error
	}

	return hasAllRoles, nil
}