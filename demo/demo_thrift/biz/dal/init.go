package dal

import (
	"github.com/pikassu/Gomall/gomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/pikassu/Gomall/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
