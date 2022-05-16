package main

import (
	"douyin/routers/userRouters"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//初始化gin
	r := gin.Default()
	initRouters(r) //初始化路由
	r.Run()        // 监听 0.0.0.0:8080
}

func initRouters(r *gin.Engine) {
	userRouters.UserRoutersInit(r)
}
