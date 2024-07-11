package database

import (
	"log"

	"github.com/silasprd/sailor-location-service/location-api/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MySQLInitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/db_location?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = DB.AutoMigrate(&entity.Location{})
	if err != nil {
		log.Fatal("Failed to migrate database schema: ", err)
	}
}
