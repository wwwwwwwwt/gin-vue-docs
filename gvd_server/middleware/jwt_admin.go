/*
 * @Author: zzzzztw
 * @Date: 2023-07-09 13:01:46
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-09 13:02:14
 * @FilePath: /gvdoc/gvd_server/middleware/jwt_admin.go
 */
package middleware

import (
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

func JwtAdmin() func(*gin.Context) {
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

		if claims.RoleID != 1 {
			res.FailWithMsg("权限错误", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
