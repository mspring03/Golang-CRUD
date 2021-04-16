package middleware_test

import (
	"net/http"
	test "net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mspring03/Golang-CRUD/user/delivery/http/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCORS(t *testing.T) {
	e := gin.New()
	_ = test.NewRequest(http.MethodGet, "/", nil)
	res := test.NewRecorder()

	c, _ := gin.CreateTestContext(res)
	m := middleware.InitMiddleware()

	h := m.CORS(gin.HandlerFunc(func(c *gin.Context) error {
		return c.ContentType(http.StatusOK)
	}))


	require.NoError(t, err)
	assert.Equal(t, "*", res.Header().Get("Access-Control-Allow-Origin"))
}
