package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/gorm/logger"
)

//DSN
const DSN = "xiaoao:bytedance@tcp(1.117.27.35:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB
var err error

//初始化数据库
func init() {
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}
