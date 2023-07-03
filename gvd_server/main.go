/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 14:15:55
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 20:38:50
 * @FilePath: /gvd_server/main.go
 */
package main

import (
	"fmt"
	"gvd_server/core"
	"gvd_server/flags"
	"gvd_server/global"
	"gvd_server/routers"
)

func main() {

	//加载日志
	global.Log = new(global.LogServer)
	global.Log.Logger = core.InitLogger()
	global.Config = core.InitConfig()

	//加载mysql和redis
	global.DB = core.InitMysql()
	global.Redis = core.InitRedis()

	option := flags.Parse()
	option.Run()

	fmt.Println(global.Config)
	router := routers.Routers()
	addr := global.Config.System.Addr()
	router.Run(addr)
}
