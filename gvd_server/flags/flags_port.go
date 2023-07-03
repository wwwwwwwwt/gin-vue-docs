/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 20:40:01
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 20:53:23
 * @FilePath: /gvd_server/flags/flags_port.go
 */
package flags

import (
	"fmt"
	"gvd_server/global"
)

func Port(port int) {
	global.Config.System.Port = port
	fmt.Println("更换端口, port:", port)
}
