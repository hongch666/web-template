package model

import "gorm.io/gorm"

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
