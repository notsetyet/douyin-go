package userController

import (
	_ "douyin/core"
	"douyin/model"
	"douyin/service/jwt"
	"douyin/service/userService"
	"douyin/vo/common"
	"douyin/vo/userVo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	user := &model.User{
		Username: username,
		Password: password,
	}

	bool := userService.CheckUser(user)
	if !bool {
		c.JSON(http.StatusOK, userVo.UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "用户名或密码错误"},
		})
		return
	}
	////封装token存储user状态
	////
	////
	////
	// 根据 user.ID 生成 token
	token, _ := jwt.SetToken(c, user.ID)
	c.JSON(http.StatusOK, userVo.UserLoginResponse{
		Response: common.Response{StatusCode: 0, StatusMsg: "登陆成功"},
		UserId:   user.ID,
		Token:    token,
	})
}
