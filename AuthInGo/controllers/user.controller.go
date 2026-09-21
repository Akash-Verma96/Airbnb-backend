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


func (uc *UserController) Ping(w http.ResponseWriter, r *http.Request){

	utils.WriteJsonSuccessResponse(w,http.StatusOK, "Service Health OK! check Completed", "NULL");
}


func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	
	payload := r.Context().Value("payload-key").(dto.CreateUserDTO)

	data, err := uc.UserService.Create(&payload)

	if err != nil || data.JwtToken == "" {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError, "Failed to Login!", err)
		return
	}


	utils.WriteJsonSuccessResponse(w,http.StatusOK, "User Created Successfully", data)
}





func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	
	
	payload := r.Context().Value("payload-key").(dto.LoginUserDTO)


	data, err := uc.UserService.LoginUser(&payload)


	if err != nil || data.JwtToken == "" {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError, "Failed to Login!", err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK, "User Logged In succesfully!", data)
}




func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("userId").(int64)

	if userId == 0 {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User ID is missing in the request context", nil)
		return
	}

	
	user, err := uc.UserService.GetUserById(userId)

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