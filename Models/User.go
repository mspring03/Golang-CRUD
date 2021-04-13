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

func Migrate(db *gorm.DB) {
	user := db.AutoMigrate(&User{})
	fmt.Println(user)
}