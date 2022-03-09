package lib

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(env Env) Database {
	// DBUsername := env.DbUsername
	// DBPassword := env.DbPassword
	// DBHost := env.DbHost
	// DBPort := env.DbPort
	// DBName := env.DbName

	url := "root:@tcp(127.0.0.1:3306)/kinder?charset=utf8mb4&parseTime=True&loc=Local"

	// url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUsername, DBPassword, DBHost, DBPort, DBName)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection Error")
	} else {
		fmt.Println("Connection Success")
	}

	return Database{
		DB: db,
	}
}
