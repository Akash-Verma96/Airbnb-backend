package controllers

import (
	dto "AuthInGo/dto"
	"AuthInGo/services"
	utils "AuthInGo/utils"
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
	
	payload := r.Context().Value("payload-key").(dto.CreateUserDTO)

	uc.UserService.Create(&payload)


	utils.WriteJsonSuccessResponse(w,http.StatusOK, "User Created Successfully", "User Created!")
}





func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	
	
	payload := r.Context().Value("payload-key").(dto.LoginUserDTO)

	jwtToken, err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError, "Failed to Login!", err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK, "User Logged In succesfully!", jwtToken)
}




func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {

	var payload struct {
		Id int64 `json:"id"`
	}

	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Bad Input", jsonErr)
		return
	}
	
	user, err := uc.UserService.GetUserById(payload.Id)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error while fetching the user!", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK,"User Fetched Successful", user)
}


func (uc *UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {


	users := uc.UserService.GetAll()


	utils.WriteJsonSuccessResponse(w,http.StatusOK,"All Profile Fetched Successfully!", users)
}



func (us *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	
	var payload struct {
		Id int64 `json:"id"`
	}

	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Bad Input", jsonErr)
		return
	}

	us.UserService.DeleteById(payload.Id)


	utils.WriteJsonSuccessResponse(w,http.StatusOK, "User Deleted Successfully!", "nil")
}