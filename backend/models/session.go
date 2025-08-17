package models

import "time"

type Session struct {
	ID        uint      `gorm:"primaryKey"`
	LectureID uint      `gorm:"not null"`
	Date      time.Time `gorm:"type:date;not null"` // Stores only date (YYYY-MM-DD)
	Status    string

	Lecture Lecture `gorm:"foreignKey:LectureID"`
}
