package lib

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(env Env) Database {
	DBUsername := env.DbUsername
	DBPassword := env.DbPassword
	DBHost := env.DbHost
	DBPort := env.DbPort
	DBName := env.DbName

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DBHost, DBUsername, DBPassword, DBName, DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection Error")
	} else {
		fmt.Println("Connection Success")
	}

	return Database{
		DB: db,
	}
}
