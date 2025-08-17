package config

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

var DB *gorm.DB

func ConnectDB() {
    user := os.Getenv("PG_USER")
    host := os.Getenv("PG_HOST")
    port := os.Getenv("PG_PORT")
    password := os.Getenv("PG_PASSWORD")
    dbname := os.Getenv("PG_DATABASE")
	sslmode := os.Getenv("PG_SSLMODE")


    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=10",
        host, port, user, password, dbname, sslmode,
    )

    log.Printf("Connecting to database at %s:%s with SSL mode=%s", host, port, sslmode)

    // Create pgx Config from DSN
    config, err := pgx.ParseConfig(dsn)
    if err != nil {
        log.Fatalf("Unable to parse DSN: %v", err)
    }

    // ✅ Force IPv4 by overriding DialFunc
    dialer := &net.Dialer{
        Timeout:   10 * time.Second,
        KeepAlive: 10 * time.Second,
    }
    config.DialFunc = func(ctx context.Context, network, addr string) (net.Conn, error) {
        return dialer.DialContext(ctx, "tcp4", addr) // force IPv4
    }

    // Open sql.DB from pgx Config
    sqlDB := stdlib.OpenDB(*config)

    db, err := gorm.Open(postgres.New(postgres.Config{
        Conn: sqlDB,
    }), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    DB = db
    log.Println("✅ Successfully connected to database over IPv4")
}
