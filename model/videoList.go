package model

import (
	"douyin/core"
	"gorm.io/gorm"
)

// Auth 是对 User 的信息进行截取
type Auth struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
}

// VideoList 点赞列表的返回结构体，包含 Video 和 Auth
type VideoList struct {
	Video Video `json:"video"`
	Auth  Auth  `json:"auth"`
}

func (fuv FavoriteUserVideo) ListVideo(db *gorm.DB) ([]*VideoList, error) {
	// TODO: 这里我先查询 video 表，再到 user 表。三表查询有点复杂，可能要用到 Preload

	// 首先获取视频信息
	videos, err := fuv.List(core.DB)
	if err != nil {
		return nil, err
	}
	ret := make([]*VideoList, len(videos))

	// 遍历指针数组获取每个视频的 UserId，去 users 表中获取 auth 的信息
	//var authID uint
	//var auth User
	for idx, val := range videos {
		authID := val.UserId
		user := User{Model: gorm.Model{ID: authID}}
		auth, err := user.List(core.DB) // user 表查询
		if err != nil {
			return nil, err
		}
		// 将 video 和 user 表的查询结果整合到 VideoList 中
		ret[idx] = &VideoList{
			Video: Video{
				Model:         gorm.Model{ID: val.ID},
				PlayUrl:       val.PlayUrl,
				CoverUrl:      val.CoverUrl,
				FavoriteCount: val.FavoriteCount,
				CommentCount:  val.CommentCount,
			},
			Auth: Auth{
				ID:            auth.ID,
				Name:          auth.Username,
				FollowCount:   auth.FollowCount,
				FollowerCount: auth.FollowerCount,
			},
		}
	}
	return ret, nil
}
