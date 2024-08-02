package dal

import (
	"cart/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
