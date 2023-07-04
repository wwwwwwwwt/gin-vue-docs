/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 20:49:14
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 21:30:01
 * @FilePath: /gvd_server/testdata/parse_jwt.go
 */
package main

import (
	"fmt"
	"gvd_server/core"
	"gvd_server/global"
	"gvd_server/utils/jwts"
)

func main() {
	global.Log = new(global.LogServer)
	global.Log.Logger = core.InitLogger()
	global.Config = core.InitConfig()

	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: "ztw",
		RoleID:   1,
		UserId:   2,
	})

	claims, err := jwts.ParseToken(token)
	fmt.Println(claims, err)
}
