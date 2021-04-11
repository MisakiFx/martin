package redis

import (
	"github.com/MisakiFx/martin/martin/pkg/tools"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func Init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		tools.GetLogger().Errorf("Init redis client error : %v", err)
		panic(err)
	}
	tools.GetLogger().Infof("Init redis client ping : %v", pong)
	redisClient = client
}

func GetRedisClient() *redis.Client {
	return redisClient
}
