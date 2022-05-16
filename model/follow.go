package model

//关注表  实际上为user to user 中间表 用于关注和粉丝
type Follow struct {
	UserId   uint `json:"user_id"`
	ToUserId uint `json:"to_user_id"` //UserId对应关注的user
}
