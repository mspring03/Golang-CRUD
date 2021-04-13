package Delivery

import "github.com/gin-gonic/gin"

type userDelivery struct {
	uu userUsecase
}

type userUsecase interface {
	Signup(c *gin.Context)
}

func NewUserDelivery(uu userUsecase) *userDelivery {
	return &userDelivery{uu}
}

func (ud *userDelivery) Routing(router *gin.RouterGroup) {
	router.POST("/signup", ud.uu.Signup)
}

