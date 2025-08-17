package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	user := os.Getenv("PG_USER")
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DATABASE")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db
}
