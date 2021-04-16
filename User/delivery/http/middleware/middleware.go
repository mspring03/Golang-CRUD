package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

type GoMiddleware struct {

}

func (m *GoMiddleware) CORS(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		next(c)
	}
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