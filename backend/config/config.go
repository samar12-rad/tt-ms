package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
    user := os.Getenv("PG_USER")
    host := os.Getenv("PG_HOST")
    port := os.Getenv("PG_PORT")
    password := os.Getenv("PG_PASSWORD")
    dbname := os.Getenv("PG_DATABASE")

    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=require connect_timeout=10",
        host, port, user, password, dbname,
    )

    log.Printf("Connecting to database at %s with SSL mode: require", host)

    db, err := gorm.Open(postgres.New(postgres.Config{
        DSN:                  dsn,
        PreferSimpleProtocol: true,
    }), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    DB = db
    log.Println("Successfully connected to database")
}
