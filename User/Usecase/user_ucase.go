package Usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/mspring03/Golang-CRUD/domain"
	"github.com/mspring03/Golang-CRUD/domain/JWT"
	"net/http"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) *userUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) Signup(c *gin.Context) {
	resp := gin.H{"state": 0, "code": 0, "message": ""}

	reqBody := new(domain.SignupRequestBody)
	err := c.Bind(reqBody)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := uu.userRepo.FindOneId(reqBody.Id)
	if user.ID != "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	uu.userRepo.CreateUser(reqBody.Id, reqBody.Password, reqBody.Age)
	token, err := JWT.CreateToken(reqBody.Id)
	if err != nil {

	}

	resp["state"] = http.StatusCreated
	resp["message"] = "User Account Creation Successful"
	resp["token"] = token
	c.JSON(http.StatusCreated, resp)
}

