package main // import "github.com/mspring03/Golang-CRUD"

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"

	"github.com/mspring03/Golang-CURD/Models"
	"github.com/mspring03/Golang-CURD/Repository"
	"github.com/mspring03/Golang-CURD/Usecase"
)

func main() {
	r := gin.Default()

	dsn := os.Getenv("DatabaseUrl")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

	}

	Models.UserMigrate(db)

	ur := Repository.UserRepo(db)

	uu := Usecase.NewUserUsecase(r, ur)





	_ = r.Run(":8080")
}