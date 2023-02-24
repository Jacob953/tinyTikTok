package apiserver

import (
	"github.com/gin-gonic/gin"

	c "github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/controller/comment"
	f "github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/controller/favorite"
	. "github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/controller/feed"
	p "github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/controller/publish"
	u "github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/controller/user"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/apiserver/store/mysql"
	"github.com/CSU-Apple-Lab/tinyTikTok/internal/pkg/middleware/auth"
)

func initRouter(g *gin.Engine) {
	// douyin handlers, requiring authentication
	storeIns, _ := mysql.GetMySQLFactory(nil)
	douyin := g.Group("/douyin")
	{
		// Middlewares.
		jwtStrategy, _ := newJWTAuth().(auth.OauthStrategy)

		// feed RESTful resource
		douyin.GET("/feed/", NewFeedController(storeIns).Get)

		// user RESTful resource
		user := douyin.Group("/user")
		{
			userController := u.NewUserController(storeIns)

			user.GET("/", userController.Get)
			user.POST("/register/", userController.Register)
			user.POST("/login/", jwtStrategy.LoginHandler)
		}

		// publish RESTful resource
		publish := douyin.Group("/publish")
		{
			publishController := p.NewPublishController(storeIns)

			publish.POST("/action/", publishController.Action)
			publish.GET("/list/", publishController.List)
		}

		// favorite RESTful resource
		favorite := douyin.Group("/favorite")
		{
			favoriteController := f.NewFavoriteController(storeIns)

			favorite.POST("/action/", favoriteController.Action)
			favorite.GET("/list/", favoriteController.List)
		}

		// comment RESTful resource
		comment := douyin.Group("/comment")
		{
			commentController := c.NewCommentController(storeIns)
			comment.POST("/action/", commentController.Action)
			comment.GET("/list/", commentController.List)
		}
	}
}
