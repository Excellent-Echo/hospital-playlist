package config

import (
	"fmt"
	"hospital-playlist/migration"
	"log"
	"os"

	// "log"

	// "github.com/joho/godotenv"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	// dsn := "pdPjydcly2:NnoFX8kkvm@tcp(remotemysql.com:3306)/pdPjydcly2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&migration.User{})
	db.AutoMigrate(&migration.UserDetail{})
	db.AutoMigrate(&migration.UserProfile{})
	db.AutoMigrate(&migration.Drug{})
	db.AutoMigrate(&migration.Specialist{})
	db.AutoMigrate(&migration.Booking{})
	db.AutoMigrate(&migration.Dokter{})

	return db
}
