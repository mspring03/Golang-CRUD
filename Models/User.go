package Models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Password  string
	Age       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

func UserMigrate(db *gorm.DB) *User {
	user := &User{}
	err := db.AutoMigrate(user)

	if err != nil {

	}

	return user
}