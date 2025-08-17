package migrations

import (
	"tms-server/config"
	"tms-server/models"
)

// INFO: for UP and DOWN migration: github.com/golang-migrate/migrate/v4
func Migrate() error {
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Faculty{},
		&models.Course{},
		&models.Batch{},
		&models.Subject{},
		&models.Room{},
		&models.Lecture{},
		&models.Session{},
	)
	return err
}
