package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/middleware"
)

// AutoStrategy defines authentication strategy which can automatically choose between Basic and Bearer
// according `Authorization` header.
type AutoStrategy struct {
	oauth middleware.AuthStrategy
}

var _ middleware.AuthStrategy = &AutoStrategy{}

// NewAutoStrategy create auto strategy with basic strategy and jwt strategy.
func NewAutoStrategy(oauth middleware.AuthStrategy) AutoStrategy {
	return AutoStrategy{
		oauth: oauth,
	}
}

// AuthFunc defines auto strategy as the gin authentication middleware.
func (a AutoStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		operator := middleware.AuthOperator{}

		// This supposed to be switch by Authorization
		operator.SetStrategy(a.oauth)

		operator.AuthFunc()(c)

		c.Next()
	}
}
