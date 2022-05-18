package commentService

import (
	"douyin/dao/commentDao"
	"douyin/model"
	"douyin/vo/commentVo"
	"douyin/vo/common"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
)

//查看评论

func SelectCommentList(vid uint, uid uint) (clr commentVo.CommentListResponse) {
	//uid是前端传进的用户id
	//根据video_id查询评论列表
	err, comments := commentDao.SelectByVideoIdDescByCreateTime(vid)

	//对错误进行判断
	if err == nil {
		if len(comments) == 0 { //如果没有评论
			clr.Response.StatusCode = 0
			clr.Response.StatusMsg = "暂无评论"
			//fmt.Println("comments--", cap(comments))
		} else { //有评论的话，根据查询到的数据，为和前端对应的结构体赋值
			clr.Response.StatusCode = 0
			clr.Response.StatusMsg = "success"
			CommentList := make([]commentVo.Comment, len(comments))
			for i, comment := range comments {
				CommentList[i].Id = comment.ID
				//根据UserId查询User表，返回一个User对象，然后转为commentVO的User形式
				{
					err, user := commentDao.SelectUserById(comment.UserId)
					if err == nil {
						CommentList[i].User.Id = user.ID
						CommentList[i].User.Name = user.Username
						CommentList[i].User.FollowCount = user.FollowCount
						CommentList[i].User.FollowerCount = user.FollowerCount
						{
							err, exist := commentDao.SelectByTwoId(uid, comment.UserId)
							if err == nil {
								CommentList[i].User.IsFollow = exist
							} else {
								fmt.Println(err) //TODO 先打印到控制台，后续可改为记到日志中
							}
						}
					} else {
						fmt.Println(err) //TODO 先打印到控制台，后续可改为记到日志中
					}
				}
				CommentList[i].Content = comment.Content
				CommentList[i].CreateDate = comment.CreatedAt
			}
			clr.CommentList = CommentList
		}
	} else {
		fmt.Println(err) //TODO 先打印到控制台，后续可改为记到日志中
		clr.Response.StatusCode = 1
		clr.Response.StatusMsg = "error"
	}
	return clr
}

//添加评论

func AddComment(comment *model.Comment) (resp common.Response) {
	//TODO 下面这两个操作应该是一个原子操作，需要手动开启关闭事务
	err1 := commentDao.AddComment(comment)
	err2, _ := commentDao.UpdateCommentCount(1, comment.VideoId)
	if err1 == nil && err2 == nil {
		resp.StatusCode = 0
		resp.StatusMsg = "success"
	} else {
		fmt.Println(err1) //TODO 先打印到控制台，后续可改为记到日志中
		fmt.Println(err2)
		resp.StatusCode = 1
		resp.StatusMsg = "error"
	}
	return resp
}

//删除评论

func DelComment(id uint) (resp common.Response) {
	//TODO 下面这三个个操作应该是一个原子操作，需要手动开启关闭事务
	err, comment := commentDao.SelectCommentById(id)
	err1 := commentDao.DelCommentById(id)
	err2, _ := commentDao.UpdateCommentCount(-1, comment.VideoId)
	if err == nil && err1 == nil && err2 == nil {
		resp.StatusCode = 0
		resp.StatusMsg = "success"
	} else {
		fmt.Println(err1) //TODO 先打印到控制台，后续可改为记到日志中
		fmt.Println(err2)
		resp.StatusCode = 1
		resp.StatusMsg = "error"
	}
	return resp
}

//一个工具函数，用来处理编码问题,将传进来的context转为latin1编码

func Convert1(src string) (string, error) {
	gbk, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(src))
	if err != nil {
		return "", err
	}
	latin1, err := charmap.ISO8859_1.NewDecoder().Bytes(gbk)
	if err != nil {
		return "", err
	}
	return string(latin1), nil
}

//工具函数，处理编码问题，将Latin1编码解码

func Convert2(src string) (string, error) {
	iso, err := charmap.ISO8859_1.NewEncoder().Bytes([]byte(src))
	if err != nil {
		return "", err
	}
	gbk, err := simplifiedchinese.GBK.NewDecoder().Bytes(iso)
	if err != nil {
		return "", err
	}
	return string(gbk), nil
}
