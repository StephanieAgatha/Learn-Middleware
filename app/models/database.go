package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database" + err.Error())
		return nil, err
	}

	db.AutoMigrate(&User{})
	DB = db
	return db, nil
}
