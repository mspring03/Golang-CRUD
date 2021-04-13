package Usecase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mspring03/Golang-CRUD/domain"
	"net/http"
	"strconv"
)

type userUsecase struct {
	ur domain.UserRepo
}

func NewUserUsecase(ur domain.UserRepo) *userUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) Signup(c *gin.Context) {
	reqBody := new(domain.SignupRequestBody)
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

