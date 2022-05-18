package commentDao

import (
	"douyin/core"
	"douyin/model"
)

/*
	对users表的操作
*/

/*
	查询评论列表的时候需要根据userid查询user对象
*/

//根据用户id查询用户

func SelectUserById(id uint) (err error, user model.User) {
	err = core.DB.Where("id=?", id).Find(&user).Error
	return err, user
}
