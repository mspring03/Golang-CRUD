package main // import "github.com/mspring03/Golang-CRUD"

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	mysql2 "github.com/mspring03/Golang-CRUD/User/repository/mysql"
	"github.com/mspring03/Golang-CRUD/user/delivery/http"
	"log"
	"os"

	"github.com/mspring03/Golang-CRUD/User/usecase"
)


func main() {
	r := gin.Default()

	dsn := os.Getenv("DatabaseUrl")
	db, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()


	ur := mysql2.UserRepo(db)
	uu := usecase.NewUserUsecase(ur)
	http.NewUserHandler(uu, r.Group("/user"))


	_ = r.Run(":8080")
}