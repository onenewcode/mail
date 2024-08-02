package dal

import (
	"payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
