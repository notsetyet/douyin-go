package favoriteController

import (
	"douyin/dao/favoriteDao"
	"douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	// TODO: 如果想要数值类型的参数，gin 需要转化吗？
	// userID := c.PostForm("user_id")
	videoID, err := strconv.ParseInt(c.PostForm("video_id"), 10, 64)
	actionType, err := strconv.ParseInt(c.PostForm("action_type"), 10, 32) // 1-点赞	2-取消点赞

	// TODO:token鉴权

	if err = favoriteDao.UpdateFavorite(uint(videoID), uint32(actionType)); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error()})
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	}
}

// FavoriteList 点赞列表的 handler 函数
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, favoriteListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: demoVideo,
	})
}
