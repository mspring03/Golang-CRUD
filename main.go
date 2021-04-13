package main // import "github.com/mspring03/Golang-CRUD"

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/mspring03/Golang-CRUD/Models/User.go"
	"os"
)

func main() {
	r := gin.Default()

	dsn := os.Getenv("DatabaseUrl")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)

	r.Run(":8080")
}