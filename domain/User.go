package domain

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type User struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Age      uint8  `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}


type UserUsecase interface {
	Signup(c context.Context) (resp gin.H, err error)
}

type UserRepository interface {
	CreateUser(id string, pw string, age uint8)
	IdConflictCheck(ctx context.Context, a string) (res *User, err error)
}