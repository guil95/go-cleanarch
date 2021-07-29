package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connect() *gorm.DB{

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"))

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
		return nil
	}

	return connection
}
