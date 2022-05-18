package model

//用户video点赞中间表    user对video的点赞

type FavoriteUserVideo struct {
	UserId  uint `json:"user_id"`
	VideoId uint `json:"video_id"`
}
