/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 20:04:18
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 21:20:33
 * @FilePath: /gvd_server/utils/jwts/enter.go
 */
package jwts

import "github.com/dgrijalva/jwt-go/v4"

type JwtPayLoad struct {
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID"`
	UserId   uint   `json:"userID"`
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
