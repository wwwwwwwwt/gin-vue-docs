/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 22:06:09
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-05 22:20:20
 * @FilePath: /gvd_server/api/user_api/user_create.go
 */
package userapi

import (
	"gvd_server/service/common/res"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	UserName string `json:"userName" bidning:"required"` //用户名
	Password string `json:"password"`                    //密码
	NickName string `json:"nickName"`                    //昵称
	RoleID   uint   `json:"roleID" binding:"required"`   //角色id
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)

	if err != nil {
		res.FailWithMsg("系统校验失败", c)
		return
	}

	res.OKWithMsg("成功啦", c)
	return
}
