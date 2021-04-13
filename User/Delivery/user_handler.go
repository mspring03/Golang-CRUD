package Delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/mspring03/Golang-CRUD/domain"
)

type userDelivery struct {
	userUcase domain.UserUsecase
}

func NewUserDelivery(uu domain.UserUsecase) *userDelivery {
	return &userDelivery{uu}
}

func (ud *userDelivery) Routing(router *gin.RouterGroup) {
	router.POST("/signup", ud.userUcase.Signup)
}

