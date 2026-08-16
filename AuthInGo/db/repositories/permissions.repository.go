package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type PermissionRepository interface {
	CreatePermission(name string, description string, resource string, action string) (*models.Permission, error)
	GetPermissionById(id int64) (*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission, error)
	GetAllPermissions() ([]*models.Permission, error)
	DeletePermissionById(id int64) error
	UpdatePermission(id int64, name string, description string, resource string, action string) (*models.Permission, error)
}

type PermissionRepositoryImpl struct {
	db *sql.DB
}

func NewPermissionRepository(_db *sql.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db: _db,
	}
}

func (p *PermissionRepositoryImpl) CreatePermission(name string, description string, resource string, action string) (*models.Permission, error) {
	query := "INSERT INTO permissions (name, description, resource, action) VALUES (? ? ? ? ?)"

	row , err := p.db.Exec(query,name,description,resource,action);

	if err != nil {
		fmt.Println("Error while executing error", err)
		return nil,err
	}

	rowsAffected , affectedErr := row.RowsAffected()

	if affectedErr != nil {
			fmt.Println("Error while checking affected row!", affectedErr)
			return nil,affectedErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not created")
		return nil,nil
	}

	id, err := row.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.Permission{
		Id:          id,
		Name:        name,
		Description: description,
		Resource:   resource,
		Action:     action,
		CreatedAt:   "", // Will be set by the database
		UpdatedAt:   "", // Will be set by the database
	}, nil
}

func (p *PermissionRepositoryImpl) GetPermissionById(id int64) (*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE id = ?"

	result  := p.db.QueryRow(query,id)

	permission := &models.Permission{}

	err := result.Scan(&permission.Id, &permission.Name, &permission.Description,&permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found!", err)
			return nil,err
		} else {
			fmt.Println("Error while fetching permission", err)
			return nil,err
		}
	}

	fmt.Println("permission fetched Successfully!")

	return permission,nil
}

func (p *PermissionRepositoryImpl) GetPermissionByName(name string) (*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE name = ?"

	result := p.db.QueryRow(query,name)

	permission := &models.Permission{}

	err := result.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found!", err)
			return nil,err
		} else {
			fmt.Println("Error while fetching permission", err)
			return nil,err
		}
	}

	fmt.Println("permission fetched Successfully!")

	return permission,nil
}

func (p *PermissionRepositoryImpl) GetAllPermissions() ([]*models.Permission, error) {
	query := "SELECT * FROM permissions"

	rows, err := p.db.Query(query)

	if err != nil {
		fmt.Println("Error while fetching Permissions", err)
		return nil,err
	}

	var permissions []*models.Permission

	// Loop through the result set row by row
	for rows.Next() {
		var permission = &models.Permission{}
		// Scan the database columns into the fields of the struct
		err := rows.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
		if err != nil {
			return nil, err
		}
		// Append each permission to your slice
		permissions = append(permissions, permission)
	}

	if err = rows.Err(); err != nil {
		return nil,err
	}

	fmt.Println("Permissions fetched successfully!")
	return permissions, nil
}

func (p *PermissionRepositoryImpl) DeletePermissionById(id int64) error {
	fmt.Println("Deleting the Permission by ID")

	query := "DELETE FROM permissions WHERE id = ?"

	result, err := p.db.Exec(query,id)

	if err != nil {
		fmt.Println("error while executing the query", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	fmt.Println("User Deleted Successfully!")

	return nil
}

func (p *PermissionRepositoryImpl) UpdatePermission(id int64, name string, description string, resource string, action string) (*models.Permission, error) {
	query := "UPDATE permissions SET name = ?, description = ?, resource = ?, action = ?, updated_at = NOW() WHERE id = ?"
	
	_, err := p.db.Exec(query, name, description, resource, action, id)
	if err != nil {
		return nil, err
	}

	return &models.Permission{
		Id:          id,
		Name:        name,
		Description: description,
		Resource: resource,
		Action: action,
		CreatedAt:   "", // Will be set by the database
		UpdatedAt:   "", // Will be set by the database
	}, nil
	}