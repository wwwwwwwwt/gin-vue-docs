/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 22:18:23
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 22:19:11
 * @FilePath: /gvd_server/api/enter.go
 */
package api

import userapi "gvd_server/api/user_api"

type Api struct {
	UserApi userapi.UserApi
}

var App = new(Api)
