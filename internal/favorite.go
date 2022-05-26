package internal

import "douyin/model"

// favorite 路由请求的参数校验结构体

type FavoriteActionRequest struct {
	UserID     uint `json:"user_id" binding:"required,gte=0"`
	VideoID    uint `json:"video_id" binding:"required,gte=0"`
	ActionType int8 `json:"action_type" binding:"required,oneof=1 2"`
}

// GET 参数绑定需要修饰 form，其值为url中的变量名
type FavoriteListRequest struct {
	UserID uint `form:"user_id" binding:"required,gte=0"`
}

type FavoriteListResponse struct {
	model.Response
	VideoList []*model.VideoList `json:"video_list"`
}
