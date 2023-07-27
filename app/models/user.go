package models

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not_null"`
}
