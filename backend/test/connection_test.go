package main

import (
	"log"
	"os"
	"tms-server/config"
)

func init() {
	// Change to parent directory so we can load the .env file
	if err := os.Chdir(".."); err != nil {
		log.Fatal("Failed to change directory:", err)
	}
	config.LoadEnvVariables()
}

func main() {
	log.Println("🔍 Testing Supabase connection with transaction pooling...")
	
	// Connect to database
	config.ConnectDB()
	
	// Test the connection
	if config.DB == nil {
		log.Fatal("❌ Database connection failed")
	}
	
	// Get pool statistics
	stats := config.GetPoolStats()
	log.Printf("📊 Pool Statistics: %+v", stats)
	
	// Test a simple query
	var result struct {
		Version string
	}
	
	if err := config.DB.Raw("SELECT version() as version").Scan(&result).Error; err != nil {
		log.Printf("❌ Query test failed: %v", err)
	} else {
		log.Printf("✅ Connection successful!")
		log.Printf("📍 PostgreSQL version: %s", result.Version)
	}
	
	// Test pool stats after query
	statsAfter := config.GetPoolStats()
	log.Printf("📊 Pool Statistics after query: %+v", statsAfter)
	
	// Clean up
	config.ClosePool()
	log.Println("🎉 Connection test completed successfully!")
}
