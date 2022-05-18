package commentVo

import (
	"douyin/vo/common"
	"time"
)

type User struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
type Comment struct {
	Id         uint      `json:"id"`
	User       User      `json:"user"`
	Content    string    `json:"content"`
	CreateDate time.Time `json:"create_date"`
}
type CommentListResponse struct {
	common.Response
	CommentList []Comment `json:"comment_list,omitempty"`
}
