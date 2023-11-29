package config

import (
	"fmt"
	"go-restful-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Asia%vJakarta",
		ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_DATABASE, "%2F",
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connected database")
	}

	err = db.AutoMigrate(&models.Author{}, &models.Book{})
	if err != nil {
		log.Fatal("Migrate database failed.")
	}

	DB = db
	log.Println("Database connected")
}
