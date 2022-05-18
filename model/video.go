package model

import "gorm.io/gorm"

//视频

type Video struct {
	gorm.Model
	UserId        uint   `json:"user_id"`        //视频对应的用户
	PlayUrl       string `json:"play_url"`       //视频的地址
	CoverUrl      string `json:"cover_url"`      //封面地址
	FavoriteCount int64  `json:"favorite_count"` //点赞数
	CommentCount  int64  `json:"comment_count"`  //评论数
}
