package db

import (
	"sync"
	"log"
	"os"
	"fmt"

	"github.com/joho/godotenv"
    "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
    once.Do(func() {
        initializeDB()
    })
    return db
}

func initializeDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
        log.Fatalf("One or more environment variables are missing")
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

    connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    db = connection
}
