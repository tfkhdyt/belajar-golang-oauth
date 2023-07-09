package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/tfkhdyt/belajar-golang-oauth/internal/config/db"
	"github.com/tfkhdyt/belajar-golang-oauth/internal/domain/user"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open(postgres.Open(db.PostgresDatabaseUrl))
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if err := DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	log.Println("Connected to DB")
}
