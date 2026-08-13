package db

import (
	"AuthInGo/models"
	"AuthInGo/utils"
	"database/sql"
	env "AuthInGo/config/env"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// added Interface
type UserRepository interface {
	Create(name string, email string, hashedPassword string) error
	LoginUser(email string, password string) (string, error)
	GetById(id int64) (*models.User, error) 
	GetByEmail(email string) (*models.User, error)
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

func (ur *UserRepositoryImpl) Create(name string, email string, hashedPassword string) error {
	fmt.Println(("Creating the user!"))

	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	row, err := ur.db.Exec(query, name, email, hashedPassword)

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





func (ur *UserRepositoryImpl) LoginUser(email string, password string) (string, error) {
	fmt.Println("Loging the user...")
	// db call
	user, err := ur.GetByEmail(email)

	if err != nil {
		fmt.Println("User not found",err)
		return "", nil
	}

	if user == nil {
		fmt.Println("No User found with this email!")
		return "", nil
	}

	passwordMatch  := utils.MatchPassword(password,user.Password)

	if passwordMatch != true {
		fmt.Println("Invalid credential", passwordMatch)
	}

	fmt.Println("User Logged in Successful")

	payload := jwt.MapClaims{
		"email": user.Email,
		"id": user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	// Sign and get the complete encoded token as a string using the secret
	jwtToken, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	return jwtToken, nil
}





func (ur *UserRepositoryImpl) GetByEmail(email string) (*models.User, error){
	fmt.Println("Fetching user by email")

	query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE email = ?"

	row := ur.db.QueryRow(query,email)

	user := &models.User{}

	err := row.Scan(&user.Id,&user.Username,&user.Email,&user.Password,&user.CreatedAt,&user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No User found with this email!")
			return nil, err
		} else {
			fmt.Println("Error while fetching the user", err)
			return nil, err
		}
	}

	fmt.Println("User found Successful")

	return user, nil
}


func (ur *UserRepositoryImpl) GetById(id int64) (*models.User,error) {
	fmt.Println("Fetching user by id")
	// 1. Prepare the query
	query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = ?"

	row := ur.db.QueryRow(query,id)

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

	

	var users []*models.User

	// Loop through the result set row by row
	for rows.Next() {
		var user = &models.User{}
		//  Scan the database columns into the fields of the struct
		err := rows.Scan(&user.Id,&user.Username,&user.Email,&user.Password,&user.CreatedAt,&user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		// Append each user to your slice
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

	_, err := ur.db.Exec(query,id)

	if err != nil {
		fmt.Println("error while executing the query", err)
		return err
	}

	fmt.Println("User Deleted Successfully!")

	return nil
}