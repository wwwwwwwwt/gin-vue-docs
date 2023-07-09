/*
 * @Author: zzzzztw
 * @Date: 2023-07-09 12:42:39
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-09 12:52:24
 * @FilePath: /gvdoc/gvd_server/middleware/jwt_auth.go
 */
package middleware

import (
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

func JwtAuth() func(*gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMsg("未携带token", c)
			c.Abort()
			return
		}

		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMsg("token 错误", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
