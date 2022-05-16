package userService

import (
	"douyin/dao/userDao"
	"douyin/model"
)

func CheckUser(user *model.User) bool {
	return userDao.CheckUser(user)
}

func CheckUserExists(username string) bool {
	return userDao.CheckUserExists(username)
}

func AddUser(user *model.User) bool {
	return userDao.AddUser(user)
}
