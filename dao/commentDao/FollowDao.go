package commentDao

import (
	"douyin/core"
	"douyin/model"
)

/*
	对follows表的操作
*/

/*
	当查询评论列表时，User对应的IsFollow字段查询该表获得
*/

//根据两个id查询

func SelectByTwoId(uid uint, tui uint) (err error, exist bool) {
	//uid是前端传来的用户id，tui是评论对应的用户id
	follow := model.Follow{}
	res := core.DB.Where("user_id=? AND to_user_id=?", uid, tui).Find(&follow)
	err = res.Error
	exist = res.RowsAffected > 0
	return err, exist
}
