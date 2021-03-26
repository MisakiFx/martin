package connection

import (
	"github.com/MisakiFx/martin/pkg/connection/mysql"
	"github.com/MisakiFx/martin/pkg/connection/redis"
)

func Init() {
	mysql.Init()
	redis.Init()
}
