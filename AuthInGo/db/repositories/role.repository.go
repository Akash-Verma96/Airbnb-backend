package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type RoleRepository interface{
	CreateRole(name string, description string) (*models.Role, error)
	GetRoleById(id int64) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	DeleteRoleById(id int64) error
	UpdateRole(id int64, name string, description string) (*models.Role, error)
}

type RoleRepositoryImpl struct{
	db *sql.DB
}

func NewRoleRepository(_db *sql.DB) RoleRepository{
	return &RoleRepositoryImpl{
		db: _db,
	}
}

func (r *RoleRepositoryImpl) CreateRole(name string, description string) (*models.Role,error){

	query := "INSERT INTO roles (name, description) VALUES (? ?)"

	row , err := r.db.Exec(query,name,description);

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

	return &models.Role{
		Id:          id,
		Name:        name,
		Description: description,
		CreatedAt:   "", // Will be set by the database
		UpdatedAt:   "", // Will be set by the database
	}, nil
}

func (r *RoleRepositoryImpl) GetRoleById(id int64) (*models.Role, error){

	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE id = ?"

	result  := r.db.QueryRow(query,id)

	role := &models.Role{}

	err := result.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found!", err)
			return nil,err
		} else {
			fmt.Println("Error while fetching role", err)
			return nil,err
		}
	}

	fmt.Println("Role fetched Successfully!")

	return role,nil
}

func (r *RoleRepositoryImpl) GetRoleByName(name string) (*models.Role, error){

	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE name = ?"

	result := r.db.QueryRow(query,name)

	role := &models.Role{}

	err := result.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found!", err)
			return nil,err
		} else {
			fmt.Println("Error while fetching role", err)
			return nil,err
		}
	}

	fmt.Println("Role fetched Successfully!")

	return role,nil
}

func (r *RoleRepositoryImpl) GetAllRoles() ([]*models.Role, error){

	query := "SELECT * FROM roles"

	rows, err := r.db.Query(query)

	if err != nil {
		fmt.Println("Error while fetching Users", err)
		return nil,err
	}

	

	var roles []*models.Role

	// Loop through the result set row by row
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

	fmt.Println("User fecthed successfully!")
	return roles, nil
}

func (r *RoleRepositoryImpl) DeleteRoleById(id int64) error{
	fmt.Println("Deleting the Role by ID")

	query := "DELETE FROM roles WHERE id = ?"

	result, err := r.db.Exec(query,id)

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

func (r *RoleRepositoryImpl) UpdateRole(id int64, name string,description string) (*models.Role,error){

	query := "UPDATE roles SET name = ?, description = ?, updated_at = NOW() WHERE id = ?"
	_, err := r.db.Exec(query, name, description, id)
	if err != nil {
		return nil, err
	}

	return &models.Role{
		Id:          id,
		Name:        name,
		Description: description,
		CreatedAt:   "", // Will be set by the database
		UpdatedAt:   "", // Will be set by the database
	}, nil
}

