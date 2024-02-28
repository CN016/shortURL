package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {

	// MySQL 配置信息
	username := "***********"         // 账号
	password := "***************"     // 密码
	host := "***********************" // 地址
	port := 3306                      // 端口
	DBname := "social_worker"         // 数据库名称
	timeout := "10s"                  // 连接超时，10秒

	//// 对密码进行URL编码
	//password = url.QueryEscape(password)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, DBname, timeout)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		println(err)
	}
}
