package usecase

import (
	"context"
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

func (uu *userUsecase) Signup(ctx context.Context, a *domain.User) (resp gin.H, err error) {
	resp = gin.H{"state": 0, "code": 0, "message": ""}

	_, err = uu.userRepo.IdConflictCheck(ctx, a.Id)
	if err != nil {
		switch err {
		case domain.ErrConflict:
			resp["state"] = http.StatusNotFound
			resp["code"] = 0
			resp["message"] = "Id Already registered"
			return
		}
	}

	uu.userRepo.CreateUser(ctx, a)
	token, err := JWT.CreateToken(a.Id)
	if err != nil {

	}

	resp["state"] = http.StatusCreated
	resp["message"] = "User Account Creation Successful"
	resp["token"] = token
	return
}