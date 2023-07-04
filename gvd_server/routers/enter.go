/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 14:17:21
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 22:29:03
 * @FilePath: /gvd_server/routers/enter.go
 */
package routers

import "github.com/gin-gonic/gin"

type RounterGroup struct {
	*gin.RouterGroup
}

func Routers() *gin.Engine {
	r := gin.Default()

	//创建了一个以api开头的api分组
	apiGroup := r.Group("api")

	//将api分组赋予了RouterGroup
	routerGroup := RounterGroup{apiGroup}
	routerGroup.UserRouter()
	return r
}
