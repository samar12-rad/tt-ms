package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"tms-server/config"
	"tms-server/migrations"
	"tms-server/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	config.ConnectDB()

	// Setup graceful shutdown
	defer config.ClosePool()

	// Handle shutdown signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shutting down gracefully...")
		config.ClosePool()
		os.Exit(0)
	}()

	if *migrate {
		if err := migrations.Migrate(); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Migrations completed. Exiting.")
		return
	}

	r := gin.Default()
	routes.RegisterRoutes(r)

	port := os.Getenv("APP_PORT")
	log.Printf("Starting server on port %s with connection pooling", port)
	r.Run(":" + port)
}
