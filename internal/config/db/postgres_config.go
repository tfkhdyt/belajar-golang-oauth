package db

import (
	"fmt"
	"os"
)

var (
	dbUsername          = os.Getenv("DB_USERNAME")
	dbPassword          = os.Getenv("DB_PASSWORD")
	dbHost              = os.Getenv("DB_HOST")
	dbPort              = os.Getenv("DB_PORT")
	dbName              = os.Getenv("DB_NAME")
	PostgresDatabaseUrl = fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbUsername, dbPassword, dbHost, dbPort, dbName,
	)
)
