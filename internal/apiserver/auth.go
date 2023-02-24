package apiserver

import (
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/iam/pkg/log"
	"github.com/spf13/viper"

	u "github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/controller/user"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/middleware"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/middleware/auth"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/model"
)

const (
	// APIServerAudience defines the value of jwt audience field.
	APIServerAudience = "iam.api.marmotedu.com"

	// APIServerIssuer defines the value of jwt issuer field.
	APIServerIssuer = "iam-apiserver"
)

func newJWTAuth() middleware.AuthStrategy {
	ginjwt, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            viper.GetString("jwt.Realm"),
		SigningAlgorithm: "HS256",
		Key:              []byte(viper.GetString("jwt.key")),
		Timeout:          viper.GetDuration("jwt.timeout"),
		MaxRefresh:       viper.GetDuration("jwt.max-refresh"),
		Authenticator:    authenticator(),
		LoginResponse:    loginResponse(),
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, nil)
		},
		RefreshResponse: refreshResponse(),
		PayloadFunc:     payloadFunc(),
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			return claims[jwt.IdentityKey]
		},
		IdentityKey:  middleware.UsernameKey,
		Authorizator: authorizator(),
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		SendCookie:    true,
		TimeFunc:      time.Now,
		// TODO: HTTPStatusMessageFunc:
	})

	return auth.NewOauthStrategy(*ginjwt)
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var r u.UserRequest
		var err error

		// support header and body both
		err = c.ShouldBindQuery(&r)
		if err != nil {
			return "", jwt.ErrFailedAuthentication
		}

		// Get the user information by the login username.
		user, err := store.Client().Users().Get(c, r.Username)
		if err != nil {
			log.Errorf("get user information failed: %s", err.Error())

			return "", jwt.ErrFailedAuthentication
		}

		// Compare the login password with the user password.
		if err := user.Compare(r.Password); err != nil {
			return "", jwt.ErrFailedAuthentication
		}

		return user, nil
	}
}

func refreshResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return func(c *gin.Context, code int, token string, expire time.Time) {
		c.JSON(http.StatusOK, gin.H{
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		})
	}
}

func loginResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return func(c *gin.Context, code int, token string, expire time.Time) {
		c.JSON(http.StatusOK, gin.H{
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		})
	}
}

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		claims := jwt.MapClaims{
			"iss": APIServerIssuer,
			"aud": APIServerAudience,
		}
		if u, ok := data.(*model.User); ok {
			claims[jwt.IdentityKey] = u.Name
			claims["sub"] = u.Name
		}

		return claims
	}
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(string); ok {
			log.L(c).Infof("user `%s` is authenticated.", v)

			return true
		}

		return false
	}
}
