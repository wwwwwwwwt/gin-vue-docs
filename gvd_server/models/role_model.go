/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 09:21:45
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 16:26:41
 * @FilePath: /gvd_server/models/role_model.go
 */
package models

type RoleModel struct {
	Model
	Title    string     `gorm:"size:16;not null;comment:角色名称" json:"title"`                                    //角色名字
	Pwd      string     `gorm:"size:64;comment:角色的密码" json:"-"`                                                //角色密码
	IsSystem bool       `gorm:"column:isSystem;comment:是否是系统角色" json:"isSystem"`                               //
	DocList  []DocModel `gorm:"many2many:role_doc_models;joinForeignKey:RoleID;JoinRederences:DocID" json:"-"` //角色拥有的文档列表
}
