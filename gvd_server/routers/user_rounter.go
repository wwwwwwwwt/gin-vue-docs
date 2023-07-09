/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 22:20:06
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-09 13:12:16
 * @FilePath: /gvdoc/gvd_server/routers/user_rounter.go
 */
package routers

import (
	"gvd_server/api"
)

func (router RounterGroup) UserRouter() {
	app := api.App.UserApi
	router.POST("user", app.UserCreateView) //创建用户
	router.POST("login", app.UserLoginView) // 用户登录
}
