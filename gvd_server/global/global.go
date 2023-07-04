/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 14:26:30
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 21:30:24
 * @FilePath: /gvd_server/global/global.go
 */
package global

import (
	"gvd_server/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 使用指针当全局变量，使其只有一个
var (
	Config *config.Config
	Log    *LogServer
	DB     *gorm.DB
	Redis  *redis.Client
)

type LogServer struct {
	*logrus.Logger
}
