package utils

import (
	"context"
	"errors"
	"log"
	"time"
	"tms-server/config"
	"tms-server/models"
)

// TODO: hash password with bcrypt
func AuthenticateUser(username, password string) (*models.User, error) {
	var user models.User

	// Use pgx pool directly to avoid GORM's prepared statement caching issues
	// This bypasses the prepared statement problem with Supabase transaction pooling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, username, password, role FROM users WHERE username = $1 LIMIT 1`
	
	row := config.Pool.QueryRow(ctx, query, username)
	
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		// Log the specific error for debugging (without exposing it to the user)
		log.Printf("Database error during authentication for user '%s': %v", username, err)
		return nil, errors.New("invalid username or password")
	}

	if user.Password != password {
		log.Printf("Password mismatch for user '%s'", username)
		return nil, errors.New("invalid username or password")
	}

	log.Printf("User '%s' authenticated successfully", username)
	return &user, nil
}
