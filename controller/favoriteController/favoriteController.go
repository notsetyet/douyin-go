package favoriteController

import (
	"douyin/dao/favUserVideoDao"
	"douyin/dao/favoriteDao"
	"douyin/internal"
	"douyin/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// FavoriteAction 点赞操作的 handler 函数
func FavoriteAction(c *gin.Context) {
	var faq internal.FavoriteActionRequest
	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	//jwtUserID, err := jwt.GetToken(c, 1)
	//if err != nil || jwtUserID != userID {
	//	c.JSON(http.StatusUnauthorized, model.Response{
	//		StatusCode: -1,
	//		StatusMsg:  err.Error() + "或 token 身份不符",
	//	})
	//	return
	//}

	// 1. 更新视频的点赞数
	video := model.Video{
		Model: gorm.Model{ID: faq.VideoID},
	}
	if err := favoriteDao.UpdateFavorite(&video, uint32(faq.ActionType)); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error()})
		return
	}

	// 2. 如果点赞操作成功，维护用户的点赞列表
	fuv := model.FavoriteUserVideo{
		UserId:  faq.UserID,
		VideoId: faq.VideoID,
	}
	if err := favUserVideoDao.UpdateFavorite(&fuv, uint32(faq.ActionType)); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{StatusCode: 0})
}

// FavoriteList 获取点赞列表的 handler 函数
func FavoriteList(c *gin.Context) {
	var flq internal.FavoriteListRequest
	// 获取 GET 参数用 ShouldBindQuery
	if err := c.ShouldBindQuery(&flq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	//jwtUserID, err := jwt.GetToken(c, 1)
	//if err != nil || jwtUserID != userID {
	//	c.JSON(http.StatusUnauthorized, model.Response{
	//		StatusCode: -1,

	//		StatusMsg:  err.Error() + "或 token 身份不符",
	//	})
	//	return
	//}

	var fuv = model.FavoriteUserVideo{UserId: flq.UserID}
	videoList, err := favUserVideoDao.ListVideo(&fuv)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, internal.FavoriteListResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "请求成功",
		},
		VideoList: videoList,
	})
}
