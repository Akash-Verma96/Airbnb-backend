package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() error{
	err := godotenv.Load()

	if(err != nil){
		fmt.Println("Error While loading env")
	}

	return err
}

func GetString(key string, fallback string) string{

	value, ok := os.LookupEnv(key)

	if(!ok){
		return fallback
	}

	return value
}

func GetInt(key string,fallback int) int{
	value, ok := os.LookupEnv(key)

	if(!ok){
		return fallback
	}

	IntVal, err := strconv.Atoi(value)

	if(err != nil){
		return fallback
	}

	return IntVal
}

func GetBool(key string,fallback bool) bool{

	value, ok := os.LookupEnv(key)

	if(!ok){
		return fallback
	}

	BoolVal, err := strconv.ParseBool(value)

	if(err != nil){
		return fallback
	}

	return  BoolVal
}