package database

import (
	"fmt"
	"log"

	"github.com/Nico164/FiberGo/app/models"
	"github.com/Nico164/FiberGo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	if err != nil {
		log.Println("Error")

	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect databasee")
	}
	fmt.Println("connection Opended to Database")

	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Student{})
	DB.AutoMigrate(&models.Teacher{})
	fmt.Println("Database Migrated")

}
