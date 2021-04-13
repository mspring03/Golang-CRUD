package Usecase

import "github.com/gin-gonic/gin"

type userUsecase struct {
	r *gin.Engine
	ur userRepo
}

type userRepo interface {
	CreateUser()
}

func NewUserUsecase(r *gin.Engine, ur userRepo) *userUsecase {
	return &userUsecase{r, ur}
}

