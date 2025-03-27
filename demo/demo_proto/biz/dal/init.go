package dal

import (
	"github.com/99n/gomall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
