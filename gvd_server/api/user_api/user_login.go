/*
 * @Author: zzzzztw
 * @Date: 2023-07-09 10:44:48
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-09 12:27:10
 * @FilePath: /gvdoc/gvd_server/api/user_api/user_kigin.go
 */
package userapi

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"
	"gvd_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	Username string `json:"userName" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

func (UserApi) UserLoginView(c *gin.Context) {
	var cr UserLoginRequest

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var user models.UserModel

	err = global.DB.Take(&user, "userName = ?", cr.Username).Error
	if err != nil {
		global.Log.Warn("用户名不存在", cr.Username)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}

	if !pwd.CheckPwd(user.Password, cr.Password) {
		global.Log.Warn("用户名不存在", cr.Username, cr.Password)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}

	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: user.NickName,
		RoleID:   user.RoleID,
		UserId:   user.ID,
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("生成token失败", c)
		return
	}

	res.OKWithData(token, c)
}
