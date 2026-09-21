package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"
)

type UserService interface {
	Create(payload *dto.CreateUserDTO) (dto.UserResult, error)
	LoginUser(payload *dto.LoginUserDTO) (dto.UserResult, error)
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

func (us * UserServiceImpl) Create(payload *dto.CreateUserDTO) (dto.UserResult,error){
	fmt.Println("Creating the User reached at Service!")
	var data dto.UserResult;

	//hashing the password
	hashedPassword, err := utils.HashPassword(payload.Password)

	if err != nil {
		fmt.Println("Error hashing password", err)
		return data,nil
	}

	User,JwtToken, err := us.UserRepository.Create(payload.Username,payload.Email,hashedPassword)

	if User == nil || JwtToken == ""{
		fmt.Println("User is nil!")
		return data,err
	}
	
	
	data.User = User
	data.JwtToken = JwtToken
	return data,nil
}

func (us *UserServiceImpl) LoginUser(payload *dto.LoginUserDTO) (dto.UserResult,error) {
	fmt.Println("Loging the user reached at service")

	User,JwtToken, err := us.UserRepository.LoginUser(payload.Email,payload.Password)

	var data dto.UserResult;
	data.User = User
	data.JwtToken = JwtToken

	return data, err
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