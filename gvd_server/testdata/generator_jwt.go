/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 20:13:24
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 20:30:05
 * @FilePath: /gvd_server/testdata/generator_jwt.go
 */
package main

import (
	"fmt"
	"gvd_server/utils/jwts"
)

func main() {
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: "ztw",
		RoleID:   2,
		UserId:   1,
	})
	fmt.Println(token, err)
}
