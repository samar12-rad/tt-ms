package models

type Lecture struct {
	ID        uint   `gorm:"primaryKey"`
	DayOfWeek string `gorm:"not null"` // e.g., Monday
	StartTime string `gorm:"not null"` // Format: "09:00"
	EndTime   string `gorm:"not null"` // Format: "10:00"

	SubjectID uint
	FacultyID uint
	BatchID   uint
	Semester  uint
	RoomID    uint

	Subject Subject
	Faculty Faculty
	Batch   Batch
	Room    Room
}
