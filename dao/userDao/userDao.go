package userDao

import (
	"douyin/core"
	"douyin/model"
	"gorm.io/gorm"
	"sync/atomic"
)

var userIdSequence = int64(1)

func CheckUser(user *model.User) bool {

	//查数据库获取user
	//
	//
	//

	if user.Username == "dwl" && user.Password == "dwl" {
		user.FollowerCount = 100
		user.FollowCount = 100
		user.ID = 1
		return true
	}
	return false
}

func CheckUserExists(username string) bool {
	err := core.DB.Where("username = ?", username).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func AddUser(user *model.User) bool {
	atomic.AddInt64(&userIdSequence, 1)
	user.ID = uint(userIdSequence)
	//	add to db
	return true
}
