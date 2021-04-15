package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)

type GoMiddleware struct {

}

func (m *GoMiddleware) SetHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0, max-age=0")
	c.Header("Last-Modified", time.Now().String())
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "-1")

	c.Next()
}

func (m *GoMiddleware) CreateToken(userId string) (string, error) {
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

func (m *GoMiddleware) VerifyToken(c *gin.Context) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		fmt.Println("token error")
		fmt.Println(err)

		c.JSON(http.StatusUnauthorized, gin.H{"state": 401, "code": 0, "message": "Invaild token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		authorized, _ := claims["authorized"].(bool)
		userId, _ := claims["user_id"].(string)

		context.WithValue(c, "authorized", authorized)
		context.WithValue(c, "user_id", userId)
	}

	c.Next()
	return
}

func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}

func extractToken(c *gin.Context) string {
	bearToken := c.GetHeader("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}