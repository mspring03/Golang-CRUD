package main // import "github.com/mspring03/Golang-CRUD"

import (
	"github.com/gin-gonic/gin"
	"github.com/mspring03/Golang-CRUD/User/Delivery"
	mysql2 "github.com/mspring03/Golang-CRUD/User/Repository/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"

	"github.com/mspring03/Golang-CRUD/Models"
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
		ur := mysql2.UserRepo(db, um)
		uu := Usecase.NewUserUsecase(ur)
		ud := Delivery.NewUserDelivery(uu)
		ud.Routing(r.Group("/user"))
	}

	_ = r.Run(":8080")
}