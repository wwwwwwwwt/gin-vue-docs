/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 09:21:45
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 10:11:41
 * @FilePath: /gvd_server/models/role_model.go
 */
package models

type RoleModel struct {
	Model
	Title    string `gorm:"size:16" json:"title"`                    //角色名字
	Pwd      string `gorm:"size:16" json:"-"`                        //角色密码
	IsSystem bool   `gorm:"column:isSyttem;size:16" json:"isSystem"` //是否系统角色
}
