package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"` // hashed password
	Role     string `gorm:"default:'faculty';not null"`
}
