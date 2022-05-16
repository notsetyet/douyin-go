package model

// common.go 定义一些基本的结构体，可以内嵌到其他结构体然后拓展

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
