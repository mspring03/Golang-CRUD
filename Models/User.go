package Models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Age       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

func UserMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	fmt.Println(err)
}