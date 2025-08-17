package models

type Subject struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Code      string `gorm:"uniqueIndex;not null"`
	CourseID  uint   `gorm:"not null"`
	Course    Course
	Faculties []Faculty `gorm:"many2many:faculty_subjects;"`
}
