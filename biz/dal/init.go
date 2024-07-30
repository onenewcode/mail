package dal

import (
	"mail/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
