/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 22:06:09
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 22:11:46
 * @FilePath: /gvd_server/api/user_api/user_create.go
 */
package userapi

import "github.com/gin-gonic/gin"

type UserCreateRequest struct {
	UserName string `json:"username" bidning:"required"` //用户名
	Password string `json:"password"`                    //密码
	NickName string `json:"nickName"`                    //昵称
	RoleID   uint   `json:"releID" binding:"required"`   //角色id
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)

	if err != nil {
		c.JSON(200, gin.H{"msg": "失败"})
		return
	}

	c.JSON(200, gin.H{"msg": "成功"})
	return
}
