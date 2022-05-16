package userVo

import "douyin/vo/common"

type UserLoginResponse struct {
	common.Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}
