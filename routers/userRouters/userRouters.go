package userRouters

import (
	"douyin/controller/favoriteController"
	"douyin/controller/userController"
	_ "douyin/core"
	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {
	defaultRoutes := r.Group("/user")
	{
		defaultRoutes.POST("/login", userController.Login)
	}

	// 扩展接口1 的点赞操作和点赞列表
	favoriteRoutes := r.Group("/favorite")
	{
		favoriteRoutes.POST("/action", favoriteController.FavoriteAction)
		favoriteRoutes.GET("/list", favoriteController.FavoriteList)
	}
}
