/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 22:20:06
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 22:26:11
 * @FilePath: /gvd_server/routers/user_rounter.go
 */
package routers

import "gvd_server/api"

func (router RounterGroup) UserRouter() {
	app := api.App.UserApi
	router.POST("user", app.UserCreateView)
}
