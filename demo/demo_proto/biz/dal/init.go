package dal

import (
	"github.com/pikassu/Gomall/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/pikassu/Gomall/gomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
