package dal

import (
	"github.com/99n/gomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/99n/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
