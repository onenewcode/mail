package model

import "mail/biz/dal/mysql"

// Init 初始化数据表
func Init() {
	mysql.DB.AutoMigrate(&User{})
}
