package commentRouters

import (
	"douyin/controller/commentController"
	"github.com/gin-gonic/gin"
)

func CommentRouterInit(r *gin.Engine) {
	apiRouter := r.Group("/douyin")

	{
		apiRouter.POST("/comment/action/", commentController.CommentAction)
		apiRouter.GET("/comment/list/", commentController.CommentList)
	}
}
