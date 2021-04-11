package connection

import (
	"github.com/MisakiFx/martin/martin/pkg/connection/mysql"
	"github.com/MisakiFx/martin/martin/pkg/connection/redis"
)

func Init() {
	mysql.Init()
	redis.Init()
}
