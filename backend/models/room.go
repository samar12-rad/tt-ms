package models

type Room struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"uniqueIndex;not null"`
	Capacity int
}
