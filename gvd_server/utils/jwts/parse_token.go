/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 20:39:14
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 21:12:09
 * @FilePath: /gvd_server/utils/jwts/parse_token.go
 */
package jwts

import (
	"gvd_server/global"

	"github.com/dgrijalva/jwt-go/v4"
)

func ParseToken(token string) (*CustomClaims, error) {
	Token, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(global.Config.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := Token.Claims.(*CustomClaims)
	if !ok {
		//数据不一致
		return nil, err
	}

	if !Token.Valid {
		//令牌无效
		return nil, err
	}
	return claims, nil
}
