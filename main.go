package main // import "github.com/mspring03/Golang-CRUD"

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"

	"github.com/mspring03/Golang-CURD/Models"
)

func main() {
	r := gin.Default()

	dsn := os.Getenv("DatabaseUrl")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


	Models.Migrate(db)
	fmt.Println(db, err)

	r.Run(":8080")
}