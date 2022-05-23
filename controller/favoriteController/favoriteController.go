package favoriteController

import (
	"douyin/dao/favUserVideoDao"
	"douyin/dao/favoriteDao"
	"douyin/model"
	"douyin/service/jwt"
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

	_, err = jwt.GetToken(c, 1)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

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

// FavoriteList 点赞列表的 handler 函数
func FavoriteList(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)

	_, err = jwt.GetToken(c, 0)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	var fuv = model.FavoriteUserVideo{UserId: uint(userID)}
	favList, err := favUserVideoDao.ListFavorite(&fuv)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// TODO: 接口的 response.video 中还有 auth 结构体
	c.JSON(http.StatusOK, favoriteListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: favList,
	})
}
