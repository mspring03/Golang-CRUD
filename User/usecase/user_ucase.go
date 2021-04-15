package usecase

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mspring03/Golang-CRUD/domain"
	"golang.org/x/crypto/bcrypt"
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
	resp = gin.H{}

	_, err = uu.userRepo.IdConflictCheck(ctx, a.Id)
	if err != nil {
		switch err {
		case domain.ErrConflict:
			resp["state"] = http.StatusConflict
			resp["code"] = 0
			resp["message"] = "Id Already registered"
			return
		}
	}

	hash, err := passwordEncoder(a.Password)
	a.Password = hash

	uu.userRepo.CreateUser(ctx, a)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp["state"] = http.StatusCreated
	resp["message"] = "User Account Creation Successful"
	return
}

func (uu *userUsecase) Signin(ctx context.Context, a *domain.User) (resp gin.H, err error) {
	resp = gin.H{}

	user, err := uu.userRepo.FindUserById(ctx, a.Id)
	if err != nil {
		switch err {
		case domain.ErrNotFound:
			resp["state"] = http.StatusNotFound
			resp["code"] = 0
			resp["message"] = "Invalid login information"
			return
		default:
			fmt.Println(err)
			return
		}
	}

	_, err = passwordCompare(user.Password, a.Password)
	if err != nil {
		switch err {
		case domain.ErrForbidden:
			resp["state"] = http.StatusForbidden
			resp["code"] = 0
			resp["message"] = "Invalid login information"
			return
		default:
			fmt.Println()
			return
		}
	}

	token, err := createToken(user.Id)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp["state"] = http.StatusOK
	resp["code"] = 0
	resp["message"] = "Login Success"
	resp["token"] = token
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

func passwordEncoder(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func passwordCompare(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return false, domain.ErrForbidden
		case bcrypt.ErrHashTooShort:
			return false, domain.ErrForbidden
		default:
			return false, err
		}
	}
	return true, nil
}