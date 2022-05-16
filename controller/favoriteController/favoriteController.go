package favoriteController

import (
	"douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// response 只定义在这个 controller 里面，因为只有这里会使用
type favoriteListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

var demoVideo = []model.Video{
	{
		UserId:        0,
		PlayUrl:       "https://www.bilibili.com/bangumi/play/ep321810?t=1",
		CoverUrl:      "https://i0.hdslb.com/bfs/archive/82d4523e2562748d050a8d8ec7ebc03fbe1a15a1.jpg",
		FavoriteCount: 999,
		CommentCount:  29,
	},
}

// FavoriteAction 点赞操作的 handler 函数
func FavoriteAction(c *gin.Context) {
	//userID := c.PostForm("user_id")
	//videoID := c.PostForm("video_id")

	// TODO: 用一个 map 记录用户是否登录
	c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusBadRequest, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// FavoriteList 点赞列表的 handler 函数
func FavoriteList(c *gin.Context) {
	// TODO: 检查用户是否登录，返回真实数据
	c.JSON(http.StatusOK, favoriteListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: demoVideo,
	})
}
