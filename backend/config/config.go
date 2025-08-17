package config

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// resolveIPv4 attempts to resolve a hostname to an IPv4 address
func resolveIPv4(hostname string) string {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		log.Printf("Warning: Could not resolve %s, using as-is: %v", hostname, err)
		return hostname
	}
	
	// Look for IPv4 addresses first
	for _, ip := range ips {
		if ip.To4() != nil {
			log.Printf("Resolved %s to IPv4: %s", hostname, ip.String())
			return ip.String()
		}
	}
	
	// If no IPv4 found, return original hostname
	log.Printf("No IPv4 address found for %s, using original hostname", hostname)
	return hostname
}

func ConnectDB() {
	user := os.Getenv("PG_USER")
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DATABASE")

	var dsn string
	
	// Determine SSL mode based on host
	if host == "localhost" || host == "127.0.0.1" {
		// Local database - disable SSL
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
		log.Println("Connecting to local database without SSL")
	} else {
		// Remote database (like Supabase) - require SSL and try to resolve to IPv4
		host = resolveIPv4(host)
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require&prefer_simple_protocol=true", user, password, host, port, dbname)
		log.Printf("Connecting to remote database with SSL: %s", host)
	}
	
	// Configure GORM with connection settings
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	
	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	
	DB = db
	log.Println("Successfully connected to database")
}
