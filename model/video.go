package model

import (
	"errors"
	"gorm.io/gorm"
)

//视频
type Video struct {
	gorm.Model
	UserId        uint   `json:"user_id"`        //视频对应的用户
	PlayUrl       string `json:"play_url"`       //视频的地址
	CoverUrl      string `json:"cover_url"`      //封面地址
	FavoriteCount int64  `json:"favorite_count"` //点赞数
	CommentCount  int64  `json:"comment_count"`  //评论数
}

// LikeVideo 对视频进行点赞或取消点赞
func (v Video) LikeVideo(db *gorm.DB, actionType uint32) error {
	var err = db.Error
	// 对操作类型进行判断，只支持 1/2
	if actionType == 1 {
		// 注意这里不能使用 db.Model() 的方式，这样会导致即使没有记录也不会返回错误
		err = db.Where("id = ?", v.ID).First(&v).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	} else if actionType == 2 {
		err = db.Where("id = ?", v.ID).First(&v).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	} else {
		err = errors.New("bad actionType")
	}

	if err != nil {
		return err
	}
	return nil
}

func (v Video) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(v).Where("id = ?", v.ID).Updates(values).Error; err != nil {
		return err
	}
	return nil
}
