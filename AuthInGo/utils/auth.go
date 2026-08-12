package utils

import (
	"fmt"
	"unsafe"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error){

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error in encrypting the password.", err)
		return "", err
	}

	// type conversion more strong than direct string(bytes)
	str := unsafe.String(&hashedPassword[0], len(hashedPassword))

	return str, nil
}

func MatchPassword(password string, hashedPassword string) bool {
	isTrue := bcrypt.CompareHashAndPassword([]byte(password),[]byte(hashedPassword));

	if isTrue == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("Password is wrong!")
		return false
	}

	return true
}