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
	Signup(ctx context.Context, a *User) (resp gin.H, err error)
	Signin(ctx context.Context, a *User) (resp gin.H, err error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (err error)
	IdConflictCheck(ctx context.Context, a string) (res *User, err error)
	FindUserById(ctx context.Context, Id string) (res *User, err error)
}

type UserMiddleware interface {
	VerifyToken(c *gin.Context)
}