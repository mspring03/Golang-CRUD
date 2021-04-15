package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mspring03/Golang-CRUD/domain"
	"net/http"

	//"net/http"
)

type userHandler struct {
	Uusecase domain.UserUsecase
}

func NewUserHandler(uu domain.UserUsecase, router *gin.RouterGroup) {
	handler := &userHandler{uu}

	router.POST("/signup", handler.Signup)
}

func (ud *userHandler) Signup(c *gin.Context) {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx := c.Request.Context()
	resp, err := ud.Uusecase.Signup(ctx, &user)
	if err != nil {
		c.JSON(resp["state"].(int), resp)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

