package dal

import (
	"user/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
