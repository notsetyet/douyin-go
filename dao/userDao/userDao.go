package userDao

import (
	"douyin/model"
)

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
