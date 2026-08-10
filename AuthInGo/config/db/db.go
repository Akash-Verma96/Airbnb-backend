package config

import (
	env "AuthInGo/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetUpDB() (*sql.DB , error) {

	cfg := mysql.NewConfig()

	cfg.User = env.GetString("DB_USER","root")
	cfg.Passwd = env.GetString("DB_PASSWORD", "root")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.Addr = env.GetString("DB_ADDR", "127.0.0.1:3306")
	cfg.DBName = env.GetString("Airbnb_auth", "Airbnb_auth")

	fmt.Println("Format DSN of sql : ",cfg.FormatDSN())

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("Error while connecting db")
		return nil, err
	}

	fmt.Println("Trying to connect DB!")
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error while connecting DB")
		return nil, pingErr
	}

	fmt.Println("DataBase Connected Successfully !")

	return db, nil
}


