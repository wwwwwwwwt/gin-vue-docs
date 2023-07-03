/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 14:26:30
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 15:13:50
 * @FilePath: /gvd_server/global/global.go
 */
package global

import (
	"gvd_server/config"

	"github.com/sirupsen/logrus"
)

// 使用指针当全局变量，使其只有一个
var (
	Config *config.Config
	Log    *LogServer
)

type LogServer struct {
	*logrus.Logger
}
