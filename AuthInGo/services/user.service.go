package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"
)

type UserService interface {
	Create(payload *dto.CreateUserDTO) error
	LoginUser(payload *dto.LoginUserDTO) (string, error)
	GetUserById(id int64) (*models.User, error)
	GetAll() ([]*models.User)
	DeleteById(id int64) error
}

type UserServiceImpl struct {
	UserRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: _userRepository,
	}
}

func (us * UserServiceImpl) Create(payload *dto.CreateUserDTO) error{
	fmt.Println("Creating the User reached at Service!")

	//hashing the password
	hashedPassword, err := utils.HashPassword(payload.Password)

	if err != nil {
		fmt.Println("Error hashing password", err)
		return nil
	}

	us.UserRepository.Create(payload.Username,payload.Email,hashedPassword)
	return nil
}

func (us *UserServiceImpl) LoginUser(payload *dto.LoginUserDTO) (string,error) {
	fmt.Println("Loging the user reached at service")
	user, err := us.UserRepository.LoginUser(payload.Email,payload.Password)

	return user, err
}

func (us * UserServiceImpl) GetUserById(id int64) (*models.User, error) {
	fmt.Println("Fetching the User reached at Service!")
	users, err := us.UserRepository.GetById(id)
	return users, err
}

func (us *UserServiceImpl) GetAll() ([]*models.User){
	fmt.Println("Fetching all users..")
	users, _ := us.UserRepository.GetAll()

	return users
}

func (us *UserServiceImpl) DeleteById(id int64) error{
	fmt.Println(("Deleting the user..."))
	us.UserRepository.DeleteById(id)
	return nil
}