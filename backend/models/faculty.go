package models

type Faculty struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"not null"`
	UserID   *uint     `gorm:"default:null"`
	User     User      `gorm:"foreignKey:UserID"`
	Subjects []Subject `gorm:"many2many:faculty_subjects;"`
}
