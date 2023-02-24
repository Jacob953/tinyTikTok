package auth

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/middleware"
)

// OauthStrategy defines jwt bearer authentication strategy.
type OauthStrategy struct {
	ginjwt.GinJWTMiddleware
}

var _ middleware.AuthStrategy = &OauthStrategy{}

// NewOauthStrategy create jwt bearer strategy with GinJWTMiddleware.
func NewOauthStrategy(gjwt ginjwt.GinJWTMiddleware) OauthStrategy {
	return OauthStrategy{gjwt}
}

// AuthFunc defines jwt bearer strategy as the gin authentication middleware.
func (j OauthStrategy) AuthFunc() gin.HandlerFunc {
	return j.MiddlewareFunc()
}
