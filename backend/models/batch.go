package models

type Batch struct {
	ID       uint   `gorm:"primaryKey"`
	Year     int    `gorm:"not null"` // e.g., 2023
	Section  string `gorm:"not null"` // e.g., A, B
	CourseID uint   `gorm:"not null"`
	Course   Course
	Lectures []Lecture
}
