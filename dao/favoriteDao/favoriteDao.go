package favoriteDao

import (
	"douyin/core"
	"douyin/model"
)

// UpdateFavorite 更新点赞数
func UpdateFavorite(video *model.Video, actionType uint32) error {
	return video.LikeVideo(core.DB, actionType)
}
