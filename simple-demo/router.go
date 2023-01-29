package main

import (
	controller2 "github.com/RaymondCode/simple-demo/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller2.Feed)
	apiRouter.GET("/user/", controller2.UserInfo)
	apiRouter.POST("/user/register/", controller2.Register)
	apiRouter.POST("/user/login/", controller2.Login)
	apiRouter.POST("/publish/action/", controller2.Publish)
	apiRouter.GET("/publish/list/", controller2.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller2.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller2.FavoriteList)
	apiRouter.POST("/comment/action/", controller2.CommentAction)
	apiRouter.GET("/comment/list/", controller2.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller2.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller2.FollowList)
	apiRouter.GET("/relation/follower/list/", controller2.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller2.FriendList)
	apiRouter.GET("/message/chat/", controller2.MessageChat)
	apiRouter.POST("/message/action/", controller2.MessageAction)
}
