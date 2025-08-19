package controllers

import (
	"net/http"
	"tms-server/config"

	"github.com/gin-gonic/gin"
)

// HealthCheck provides a basic health check endpoint
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "TMS Server is running",
	})
}

// DatabaseStats provides database connection pool statistics
func DatabaseStats(c *gin.Context) {
	stats := config.GetPoolStats()
	
	c.JSON(http.StatusOK, gin.H{
		"database_pool": stats,
	})
}
