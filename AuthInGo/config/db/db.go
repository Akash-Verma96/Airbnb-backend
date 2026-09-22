// package config

// import (
// 	env "AuthInGo/config/env"
// 	"database/sql"
// 	"fmt"

// 	"github.com/go-sql-driver/mysql"
// )

// func SetUpDB() (*sql.DB , error) {

// 	cfg := mysql.NewConfig()

// 	cfg.User = env.GetString("DB_USER","root")
// 	cfg.Passwd = env.GetString("DB_PASSWORD", "root")
// 	cfg.Net = env.GetString("DB_NET", "tcp")
// 	cfg.Addr = env.GetString("DB_ADDR", "127.0.0.1:3306")
// 	cfg.DBName = env.GetString("Airbnb_auth", "Airbnb_auth")


// 	db, err := sql.Open("mysql", cfg.FormatDSN())

// 	if err != nil {
// 		fmt.Println("Error while connecting db")
// 		return nil, err
// 	}

// 	fmt.Println("Trying to connect DB!")
// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		fmt.Println("Error while connecting DB")
// 		return nil, pingErr
// 	}

// 	fmt.Println("DataBase Connected Successfully !")

// 	return db, nil
// }





package config

import (
	env "AuthInGo/config/env"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func SetUpDB() (*sql.DB, error) {

	// Load CA certificate from Aiven (download ca.pem from Aiven console)
	rootCertPool := x509.NewCertPool()
	pem, err := os.ReadFile(env.GetString("DB_SSL_CA", "ca.pem"))
	if err != nil {
		return nil, fmt.Errorf("failed to read CA cert: %w", err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		return nil, fmt.Errorf("failed to append CA cert")
	}

	// Register custom TLS config
	tlsConfig := &tls.Config{
		RootCAs:    rootCertPool,
		MinVersion: tls.VersionTLS12,
	}
	mysql.RegisterTLSConfig("aiven", tlsConfig)

	cfg := mysql.NewConfig()
	cfg.User   = env.GetString("DB_USER", "")
	cfg.Passwd = env.GetString("DB_PASSWORD", "")
	cfg.Net    = env.GetString("DB_NET", "tcp")
	cfg.Addr   = env.GetString("DB_ADDR", "")
	cfg.DBName = env.GetString("DB_NAME", "")
	cfg.TLSConfig = "aiven"

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("error opening DB: %w", err)
	}

	fmt.Println("Trying to connect DB!")
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging DB: %w", err)
	}

	fmt.Println("Database Connected Successfully!")
	return db, nil
}