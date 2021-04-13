package main // import "github.com/mspring03/Golang-CRUD"

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"

	"github.com/mspring03/Golang-CRUD/Models"
	"github.com/mspring03/Golang-CRUD/User/Delivery"
	"github.com/mspring03/Golang-CRUD/User/Repository"
	"github.com/mspring03/Golang-CRUD/User/Usecase"
)

func main() {
	r := gin.Default()

	dsn := os.Getenv("DatabaseUrl")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

	}
	um := Models.UserMigrate(db)

	{
		ur := Repository.UserRepo(db, um)
		uu := Usecase.NewUserUsecase(ur)
		ud := Delivery.NewUserDelivery(uu)
		ud.Routing(r.Group("/user"))
	}

	_ = r.Run(":8080")
}