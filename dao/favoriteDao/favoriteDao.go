package favoriteDao

import (
	"douyin/core"
	"douyin/model"
	"gorm.io/gorm"
)

// TODO: dao层拆分
// UpdateFavorite 更新点赞数
func UpdateFavorite(id uint, actionType uint32) error {
	video := model.Video{
		Model: gorm.Model{ID: id},
	}
	return video.LikeVideo(core.DB, actionType)
}
