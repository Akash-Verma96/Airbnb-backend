package controllers

import (
	"AuthInGo/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		UserService: _userService,
	}
}


func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating the User!")
	uc.UserService.Create()
	w.Write([]byte("User Create endpoint Done"))
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetchiing the User!")
	uc.UserService.GetByIdUser()
	w.Write([]byte("User Fetch endpoint Done"))
}

func (uc *UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all Users..")
	users := uc.UserService.GetAll()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 3. Encode the slice directly into the response writer -> json.NewEncoder(w).Encode(users)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

type UserRequest struct {
	id int64
}

func (us *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting user...")
	var reqBody UserRequest
    if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

	us.UserService.DeleteById(reqBody.id)
	w.Write([]byte("User Deleted Successfully"))
}