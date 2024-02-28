package common

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")

	dsn := "host=localhost user=" + DB_USER + " password=" + DB_PASS + " dbname=" + DB_NAME + " port=" + DB_PORT + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB connect error: ", err)
	}
	fmt.Println("DB connected! ", db)

	DB = db

	return DB
}

func GetDB() *gorm.DB {
	return DB
}
