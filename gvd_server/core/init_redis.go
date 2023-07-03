/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 19:01:10
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 19:15:35
 * @FilePath: /gvd_server/core/init_redis.go
 */
package core

import (
	"context"
	"gvd_server/global"
	"time"

	"github.com/go-redis/redis"
)

func InitRedis() *redis.Client {

	redisconfig := global.Config.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisconfig.Addr(),
		Password: redisconfig.Password,
		DB:       0,
		PoolSize: redisconfig.PoolSize,
	})

	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := client.Ping().Result()
	if err != nil {
		global.Log.Fatal("redis 连接失败, err: ", err.Error())
	}
	return client
}
