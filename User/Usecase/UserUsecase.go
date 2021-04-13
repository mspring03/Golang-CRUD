package Usecase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mspring03/Golang-CRUD/Models"
	"net/http"
	"strconv"
)

type userUsecase struct {
	ur userRepo
}

type userRepo interface {
	CreateUser(id string, pw string, age uint8)
	FindOneId(id string) *Models.User
}

func NewUserUsecase(ur userRepo) *userUsecase {
	return &userUsecase{ur}
}

type signupRequestBody struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Age      uint8  `json:"age"`
}

func (uu *userUsecase) Signup(c *gin.Context) {
	reqBody := new(signupRequestBody)
	err := c.Bind(reqBody)
	fmt.Println(*reqBody)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := uu.ur.FindOneId(reqBody.Id)
	if Collision, _ := strconv.ParseBool(user.ID); Collision == true {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	uu.ur.CreateUser(reqBody.Id, reqBody.Password, reqBody.Age)

	c.JSON(http.StatusCreated, reqBody)
}

