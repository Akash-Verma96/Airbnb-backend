package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
	"fmt"
)

type UserService interface {
	GetByIdUser() error
	Create() error
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

func (us * UserServiceImpl) Create() error{
	fmt.Println("Creating the User reached at Service!")
	us.UserRepository.Create()
	return nil
}

func (us * UserServiceImpl) GetByIdUser() error{
	fmt.Println("Fetching the User reached at Service!")
	us.UserRepository.GetById()
	return nil
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