package model

import (
	"gorm.io/gorm"
)

// 用户video点赞中间表    user对video的点赞
// IsDel:	0-未删除，1-删除
type FavoriteUserVideo struct {
	UserId    uint           `json:"user_id"`
	VideoId   uint           `json:"video_id"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// FavoriteResult 联表查询的结果
type FavoriteResult struct {
	UserId uint     `json:"user_id"`
	Videos []*Video `json:"videos"`
}

// Create 当记录不存在时，创建记录
// 如果记录存在，则将 deleted_at 更新为 nil
func (fuv FavoriteUserVideo) Create(db *gorm.DB) error {
	var err = db.Error
	// 这里要调用 Unscoped 查找被软删除的记录，否则还是会创建新的
	err = db.Unscoped().Where("user_id = ? AND video_id = ?", fuv.UserId, fuv.VideoId).First(&FavoriteUserVideo{}).Update(
		"deleted_at", nil).Error
	if err != nil {
		return db.Create(&fuv).Error // 如果更新失败，说明没有这条记录，则创建
	}
	return err
}

func (fuv FavoriteUserVideo) Delete(db *gorm.DB) error {
	// gorm.DeletedAt 具有软删除
	// 直接用 db.Delete(&fuv) 会报错 "Where condition required"
	return db.Where("user_id = ? AND video_id = ?", fuv.UserId, fuv.VideoId).Delete(&FavoriteUserVideo{}).Error
}

// List 联表查询
func (fuv FavoriteUserVideo) List(db *gorm.DB) ([]*Video, error) {
	var videos []*Video
	var err error
	err = db.Table("favorite_user_videos").
		Select("videos.*").
		Joins("left join videos on videos.id = favorite_user_videos.video_id").
		Where("favorite_user_videos.user_id = ?", fuv.UserId).
		Find(&FavoriteUserVideo{}).
		Scan(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
