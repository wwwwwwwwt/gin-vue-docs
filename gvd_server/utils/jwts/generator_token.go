/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 20:07:58
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 21:28:51
 * @FilePath: /gvd_server/utils/jwts/generator_token.go
 */
package jwts

import (
	"gvd_server/global"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

func GenToken(user JwtPayLoad) (string, error) {
	claims := CustomClaims{
		JwtPayLoad: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Duration(global.Config.Jwt.Expires) * time.Hour)),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.Config.Jwt.Secret))
}
