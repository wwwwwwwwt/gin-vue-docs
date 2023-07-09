/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 22:06:09
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-09 10:40:20
 * @FilePath: /gvdoc/gvd_server/api/user_api/user_create.go
 */
package userapi

import (
	"fmt"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	UserName string `json:"userName" label:"用户名" bidning:"required"` //用户名
	Password string `json:"password" label:"密码" binding:"required"`  //密码
	NickName string `json:"nickName"`                                //昵称
	RoleID   uint   `json:"roleID" binding:"required"`               //角色id
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)

	if err != nil {
		//res.FailWithMsg("系统校验失败", c)
		//res.FailWithError(err, &cr, c)
		res.FailWithValidError(err, &cr, c)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", cr.UserName).Error
	if err == nil {
		res.FailWithMsg("用户名已经存在", c)
		return
	}
	if cr.NickName == "" {
		// 如果昵称不存在，
		var maxID uint
		global.DB.Model(models.UserModel{}).Select("max(id)").Scan(&maxID)
		cr.NickName = fmt.Sprintf("用户_%d", maxID+1)
	}
	var role models.RoleDocModel
	err = global.DB.Take(&role, cr.RoleID).Error
	if err != nil {
		res.FailWithMsg("角色不存在", c)
		return
	}
	err = global.DB.Create(&models.UserModel{
		UserName: cr.UserName,
		Password: pwd.HashPwd(cr.Password),
		NickName: cr.NickName,
		IP:       c.RemoteIP(),
		RoleID:   cr.RoleID,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("用户创建失败", c)
		return
	}
	res.OKWithMsg("成功啦", c)
	return
}
