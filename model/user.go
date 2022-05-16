package model

import (
	"gorm.io/gorm"
)

//用户表
type User struct {
	gorm.Model
	Username      string `json:"username"`       //用户名
	Password      string `json:"password"`       //密码,密文存储
	FollowCount   int64  `json:"follow_count"`   //关注数
	FollowerCount int64  `json:"follower_count"` //粉丝数
}
