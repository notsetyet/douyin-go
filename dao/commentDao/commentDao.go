package commentDao

import (
	"douyin/core"
	"douyin/model"
	"gorm.io/gorm"
)

/*
	对comments表的操作
*/

/*
	1、评论列表需要查询所有未被删除的评论，根据视频id查询
	2、评论操作有两种：
		a、添加评论，传进一个comment对象
		b、删除评论，根据评论id删除
*/

/**
查寻所有未被删除的评论，时间倒序输出
*/

func SelectByVideoIdDescByCreateTime(vid uint) (err error, comments []model.Comment) {
	err = core.DB.Order("created_at desc").Where("video_id=? AND deleted_at is null", vid).Find(&comments).Error
	//First()函数找不到record的时候，会返回ErrRecordNotFound， 而Find()则是返回nil
	return err, comments
}

/**
添加一个评论
*/

func AddComment(comment *model.Comment, tx *gorm.DB) (err error) {
	err = tx.Create(comment).Error
	return err
}

/**
删除评论
*/

func DelCommentById(id uint, tx *gorm.DB) (err error) { //软删除
	err = tx.Where("id=?", id).Delete(&model.Comment{}).Error
	return err
}

/**
查询一个评论，用于删除评论时对video表操作时需要videoid，返回一个comment
*/

func SelectCommentById(id uint) (err error, comment model.Comment) {
	err = core.DB.Where("id=?", id).Find(&comment).Error
	return
}
