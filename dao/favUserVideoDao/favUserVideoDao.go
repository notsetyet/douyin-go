package favUserVideoDao

import (
	"douyin/core"
	"douyin/model"
)

func UpdateFavorite(fuv *model.FavoriteUserVideo, actionType uint32) error {
	// 点赞
	if actionType == 1 {
		return fuv.Create(core.DB)
	} else {
		return fuv.Delete(core.DB)
	}
}

// ListFavorite 列出用户的点赞视频
func ListFavorite(fuv *model.FavoriteUserVideo) ([]*model.Video, error) {
	videos, err := fuv.List(core.DB)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
