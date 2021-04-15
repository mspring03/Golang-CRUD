package usecase

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mspring03/Golang-CRUD/domain"
	"net/http"
	"os"
	"time"
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
	if err != nil {

	}

	resp["state"] = http.StatusCreated
	resp["message"] = "User Account Creation Successful"
	return
}


func createToken(userId string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}