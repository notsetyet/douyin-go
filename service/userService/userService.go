package userService

import (
	"douyin/dao/userDao"
	"douyin/model"
)

func CheckUser(user *model.User) bool {
	return userDao.CheckUser(user)
}
