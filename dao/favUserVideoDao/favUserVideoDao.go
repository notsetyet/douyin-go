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

// ListVideo 根据用户ID查找其所有点赞视频的信息 --> videoList 结构体包含视频和视频作者(auth)的信息
func ListVideo(fuv *model.FavoriteUserVideo) ([]*model.VideoList, error) {
	videoList, err := fuv.ListVideo(core.DB)
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
