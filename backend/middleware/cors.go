package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	// Default allowed origins
	defaultOrigins := []string{
		"http://localhost:3000",
		"http://localhost:5173",
		"https://tt-ms.vercel.app",
	}
	
	// Get additional origins from environment variable
	envOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	var allowedOrigins []string
	
	if envOrigins != "" {
		// Split comma-separated origins from environment
		envOriginsList := strings.Split(envOrigins, ",")
		for _, origin := range envOriginsList {
			allowedOrigins = append(allowedOrigins, strings.TrimSpace(origin))
		}
	} else {
		// Use default origins if no environment variable is set
		allowedOrigins = defaultOrigins
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
