package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/mspring03/Golang-CRUD/Models"
)

type SignupRequestBody struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Age      uint8  `json:"age"`
}


type UserUsecase interface {
	Signup(c *gin.Context)
}

type UserRepository interface {
	CreateUser(id string, pw string, age uint8)
	FindOneId(id string) *Models.User
}