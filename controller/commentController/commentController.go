package commentController

import (
	"douyin/model"
	"douyin/service/commentService"
	"douyin/vo/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//TODO token还未使用

func CommentAction(c *gin.Context) {
	userid, err1 := strconv.ParseInt(c.Query("user_id"), 10, 64)
	//token := c.Query("token")
	videoid, err2 := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, err3 := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err1 == nil && err2 == nil && err3 == nil {
		if int(actionType) == 1 {
			comment := model.Comment{
				UserId:  uint(userid),
				VideoId: uint(videoid),
				//Content: c.Query("comment_text"),这样无法插入编码为Latin1编码的数据库e
			}
			content, err4 := commentService.Convert1(c.Query("comment_text"))
			if err4 == nil {
				comment.Content = content
				comment.CreatedAt = time.Now()
				resp := commentService.AddComment(&comment)
				c.JSON(http.StatusOK, resp)
			} else {
				fmt.Println(err4) //TODO 先打印到控制台，后续可改为记到日志中
				c.JSON(http.StatusOK, common.Response{
					StatusCode: 1,
					StatusMsg:  "error",
				})
			}
		} else if int(actionType) == 2 {
			commentId, err4 := strconv.ParseInt(c.Query("comment_id"), 10, 64)
			if err4 == nil {
				resp := commentService.DelComment(uint(commentId))
				c.JSON(http.StatusOK, resp)
			} else {
				fmt.Println(err4) //TODO 先打印到控制台，后续可改为记到日志中
				c.JSON(http.StatusOK, common.Response{
					StatusCode: 1,
					StatusMsg:  "error",
				})
			}
		} else {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 1,
				StatusMsg:  "error",
			})
		}
	} else {
		fmt.Println(err1) //TODO 先打印到控制台，后续可改为记到日志中
		fmt.Println(err2)
		fmt.Println(err3)
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "error",
		})
	}
}

func CommentList(c *gin.Context) {
	uid, err1 := strconv.ParseInt(c.Query("user_id"), 10, 64)
	//token := c.Query("token")
	vid, err2 := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err1 == nil && err2 == nil {
		commentListResp := commentService.SelectCommentList(uint(vid), uint(uid))
		for i, _ := range commentListResp.CommentList {
			convert2, err1 := commentService.Convert2(commentListResp.CommentList[i].Content)
			commentListResp.CommentList[i].Content = convert2
			if err1 != nil {
				c.JSON(http.StatusOK, common.Response{
					StatusCode: 0,
					StatusMsg:  "error",
				})
				return
			}
		}
		c.JSON(http.StatusOK, commentListResp)
	} else {
		fmt.Println(err1) //TODO 先打印到控制台，后续可改为记到日志中
		fmt.Println(err2)
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "error",
		})
	}
}
