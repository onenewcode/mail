package dal

import (
	"email/biz/dal/mysql"
	"email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
