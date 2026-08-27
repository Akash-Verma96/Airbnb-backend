package middlewares

import (
	"AuthInGo/dto"
	"AuthInGo/utils"
	"context"
	"fmt"
	"net/http"
)

func BasicMiddleare(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Inside Basic Middleware function!")
		fmt.Printf("Started %s %s \n", r.Method, r.RequestURI)

		next.ServeHTTP(w,r)
	})
}

func LoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

		var payload dto.LoginUserDTO

		if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Bad Input", jsonErr)
			return
		}

		if validationErr := utils.Validator.Struct(payload); validationErr != nil {
			utils.WriteJsonErrorResponse(w,http.StatusInternalServerError, "Invalid Input data", validationErr)
			return
		}

		// getting current context
		ctx := r.Context()

		// Derive new context with the target value
		ctx = context.WithValue(ctx,"payload-key",payload)

		//	Injecting the new context into a shallow copy of the request
		rWithCtx := r.WithContext(ctx)


		next.ServeHTTP(w,rWithCtx)
	})
}


func CreateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

		var payload dto.CreateUserDTO

		if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Bad Input", jsonErr)
			return
		}

		if validationErr := utils.Validator.Struct(payload); validationErr != nil {
			utils.WriteJsonErrorResponse(w,http.StatusInternalServerError, "Invalid Input data", validationErr)
			return
		}

		// getting current context
		ctx := r.Context()

		// Derive new context with the target value
		ctx = context.WithValue(ctx,"payload-key",payload)

		//	Injecting the new context into a shallow copy of the request
		rWithCtx := r.WithContext(ctx)


		next.ServeHTTP(w,rWithCtx)
	})
}

func CreateRoleRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

		var payload dto.CreateRoleDTO

		if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Bad Input", jsonErr)
			return
		}

		if validationErr := utils.Validator.Struct(payload); validationErr != nil {
			utils.WriteJsonErrorResponse(w,http.StatusInternalServerError, "Invalid Input data", validationErr)
			return
		}

		// getting current context
		ctx := r.Context()

		// Derive new context with the target value
		ctx = context.WithValue(ctx,"payload-key",payload)

		//	Injecting the new context into a shallow copy of the request
		rWithCtx := r.WithContext(ctx)


		next.ServeHTTP(w,rWithCtx)
	})
}

func UpdateRoleRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {

		var payload dto.UpdateRoleDTO

		if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Bad Input", jsonErr)
			return
		}

		if validationErr := utils.Validator.Struct(payload); validationErr != nil {
			utils.WriteJsonErrorResponse(w,http.StatusInternalServerError, "Invalid Input data", validationErr)
			return
		}

		// getting current context
		ctx := r.Context()

		// Derive new context with the target value
		ctx = context.WithValue(ctx,"payload-key",payload)

		//	Injecting the new context into a shallow copy of the request
		rWithCtx := r.WithContext(ctx)


		next.ServeHTTP(w,rWithCtx)
	})
}