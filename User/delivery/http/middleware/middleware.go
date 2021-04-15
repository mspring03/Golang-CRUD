package middleware

import "github.com/gin-gonic/gin"

type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) error {
		c.Header("Access-Control-Allow-Origin", "*")

		return next(c)
	}
}