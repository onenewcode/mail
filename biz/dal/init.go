package dal

import (
	"mail/biz/dal/mysql"
	"mail/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
