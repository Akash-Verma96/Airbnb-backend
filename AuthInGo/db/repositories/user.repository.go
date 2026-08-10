package db

import (
	"database/sql"
	"AuthInGo/models"
	"fmt"
)

// added Interface
type UserRepository interface {
	GetById() (*models.User, error)
	Create() error
	GetAll() ([]*models.User, error)
	DeleteById(id int64) error
}

// struct which going to implement the UserRepository
type UserRepositoryImpl struct {
	db *sql.DB
}

// constructer of UserRepository
func NewUserRepository(_db *sql.DB) UserRepository{
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (ur *UserRepositoryImpl) Create() error {
	fmt.Println(("Creating the user!"))

	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	row, err := ur.db.Exec(query, "testUser", "test@gmail.com", 1233)

	if err != nil {
		fmt.Println("Error while executing the query", err)
		return err
	}

	rowsAffected , affectedErr := row.RowsAffected()

	if affectedErr != nil {
			fmt.Println("Error while checking affected row!", affectedErr)
			return affectedErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not created")
		return nil
	}

	fmt.Println("User created Successfully!")

	return nil
}

// implimenting userRepositrory after creating creat method
func (ur *UserRepositoryImpl) GetById() (*models.User,error) {
	fmt.Println("Fetching user in UserRepository")
	// 1. Prepare the query
	query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = ?"

	row := ur.db.QueryRow(query,1)

	user := &models.User{}

	err := row.Scan(&user.Id,&user.Username,&user.Email,&user.Password,&user.CreatedAt,&user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found while fetching!")
			return nil, err
		} else {
			fmt.Println("Error while fecthing User", err)
			return nil, err
		}
	}

	fmt.Println("User fecthed successfully!", user)
	return user, nil
}

func (ur *UserRepositoryImpl) GetAll() ([]*models.User,error){
	fmt.Println("Getting All Rowas")

	query := "SELECT * FROM users"

	rows, err := ur.db.Query(query)

	if err != nil {
		fmt.Println("Error while fetching Users", err)
		return nil,err
	}

	

	// scanErr := result.Scan(&user.Id,&user.Username,&user.Email,&user.Password,&user.CreatedAt,&user.UpdatedAt)

	var users []*models.User

	// 5. Loop through the result set row by row
	for rows.Next() {
		var user = &models.User{}
		// 6. Scan the database columns into the fields of the struct
		err := rows.Scan(&user.Id,&user.Username,&user.Email,&user.Password,&user.CreatedAt,&user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		// 7. Append each user to your slice
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil,err
	}

	fmt.Println("User fecthed successfully!")
	return users, nil
}

func (ur * UserRepositoryImpl) DeleteById(id int64) error {
	fmt.Println("Deleting the user by ID")

	query := "DELETE FROM users WHERE id = ?"

	_, err := ur.db.Exec(query,2)

	if err != nil {
		fmt.Println("error while executing the query", err)
		return err
	}

	fmt.Println("User Deleted Successfully!")

	return nil
}