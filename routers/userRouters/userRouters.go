package userRouters

import (
	"douyin/controller/userController"
	_ "douyin/core"
	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {
	defaultRoutes := r.Group("/user")
	{
		defaultRoutes.POST("/login", userController.Login)
	}
}
