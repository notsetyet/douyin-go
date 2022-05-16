package model

import "gorm.io/gorm"

//评论
type Comment struct {
	gorm.Model
	UserId  uint   `json:"user_id"`  //评论对应的用户
	VideoId uint   `json:"video_id"` //评论对应的视频
	Content string `json:"content"`  //评论的信息
}
