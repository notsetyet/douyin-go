package commentDao

import (
	"douyin/core"
	"douyin/model"
)

//对videos表的comment_count的更改

//根据video_id更新

func UpdateCommentCount(num int, vid uint) (err error, ok bool) {
	//传进的num为1或者-1
	video := model.Video{}
	err1 := core.DB.Select("comment_count").Where("id=?", vid).Find(&video).Error
	if err1 != nil {
		return err1, false
	}
	err = core.DB.Model(&model.Video{}).Where("id=?", vid).Update("comment_count", int(video.CommentCount)+num).Error
	if err == nil {
		return err, true
	} else {
		return err, false
	}
}
