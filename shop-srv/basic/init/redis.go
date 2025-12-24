package init

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"shop-srv/basic/globals"
)

func InitRedis() {
	redisConfig := globals.AppConfig.Redis
	Addr := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	globals.Rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: redisConfig.Password,
		DB:       0,
	})
	err := globals.Rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("redis 连接成功")
	}
}
