package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)


var Validator *validator.Validate

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func init() {
	fmt.Println("Intializing the validator!")
	Validator = NewValidator()
}

func WriteJsonResponse(w http.ResponseWriter,status int, data any) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data) 
}

func ReadJsonBody(r *http.Request, result any) error{
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(result)
}

func WriteJsonSuccessResponse(w http.ResponseWriter, status int, message string, data any) error {
	response := map[string]any{}
	
	response["status"] = "success"
	response["message"] = message
	response["data"] = data

	return WriteJsonResponse(w, status, response)
}

func WriteJsonErrorResponse(w http.ResponseWriter, status int, message string, err error) error {
	response := map[string]any{}

	response["stauts"] = "error"
	response["message"] = message
	response["data"] = err.Error()

	return WriteJsonResponse(w, status, response)
}