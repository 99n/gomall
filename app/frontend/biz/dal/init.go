package dal

import (
	"github.com/99n/gomall/app/frontend/biz/dal/mysql"
	"github.com/99n/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
