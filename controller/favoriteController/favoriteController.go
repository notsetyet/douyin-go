package favoriteController

import (
	"douyin/dao/favUserVideoDao"
	"douyin/dao/favoriteDao"
	"douyin/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// response 只定义在这个 controller 里面，因为只有这里会使用
type favoriteListResponse struct {
	model.Response
	VideoList []*model.Video `json:"video_list"`
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
	userID, err := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	videoID, err := strconv.ParseInt(c.PostForm("video_id"), 10, 64)
	actionType, err := strconv.ParseInt(c.PostForm("action_type"), 10, 32) // 1-点赞	2-取消点赞
	// 对 actionType 进行参数校验
	if actionType != 1 && actionType != 2 {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: -1,
			StatusMsg:  "actionType must be 1 or 2",
		})
		return
	}
	// TODO:token鉴权

	// 1. 更新视频的点赞数
	video := model.Video{
		Model: gorm.Model{ID: uint(videoID)},
	}
	if err = favoriteDao.UpdateFavorite(&video, uint32(actionType)); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error()})
		return
	}

	// 2. 如果点赞操作成功，维护用户的点赞列表
	fuv := model.FavoriteUserVideo{
		UserId:  uint(userID),
		VideoId: uint(videoID),
	}
	if err = favUserVideoDao.UpdateFavorite(&fuv, uint32(actionType)); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{StatusCode: 0})
}

// TODO: 完善调用链
// FavoriteList 点赞列表的 handler 函数
func FavoriteList(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)

	var fuv = model.FavoriteUserVideo{UserId: uint(userID)}
	favList, err := favUserVideoDao.ListFavorite(&fuv)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, favoriteListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: favList,
	})

	//c.JSON(http.StatusOK, favoriteListResponse{
	//	Response: model.Response{
	//		StatusCode: 0,
	//	},
	//	VideoList: demoVideo,
	//})
}
