package databases

import (
	"fmt"
	"os"
	"week4/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Failed to Connected")
	}
	fmt.Println("Database Connection Succesfully")
	DB = db
}

func DBMigrate() {
	if err := DB.AutoMigrate(&models.Note{}); err != nil {
		panic("Database Migration Failed")
	}
	fmt.Println("Database Migration Succesfully")
}