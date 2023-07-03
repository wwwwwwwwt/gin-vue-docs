/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 14:15:55
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 15:14:09
 * @FilePath: /gvd_server/main.go
 */
package main

import (
	"fmt"
	"gvd_server/core"
	"gvd_server/global"
	"gvd_server/routers"
)

func main() {
	global.Log = new(global.LogServer)
	global.Log.Logger = core.InitLogger()
	global.Config = core.InitConfig()
	global.Log.Infof("xxx")
	fmt.Println(global.Config)
	router := routers.Routers()
	addr := global.Config.System.Addr()
	router.Run(addr)
}
