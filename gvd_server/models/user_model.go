/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 21:04:37
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 21:55:31
 * @FilePath: /gvd_server/models/user_model.go
 */
package models

type UserModel struct {
	Model
	UserName string `gorm:"column:userName;size:36;unique;not null" json:"userName"` //用户名
	Password string `gorm:"column:password;size:256" json:"password"`                //密码
	Avatar   string `gorm:"column:avatar;size:256" json:"avatar"`                    //头像
	NickName string `gorm:"column:nickName;size:36" json:"nickName"`                 //昵称
	Email    string `gorm:"column:email;size:128" json:"email"`                      //邮箱
	Token    string `gorm:"column:token;size:36" json:"token"`                       //其他平台的唯一id
	IP       string `gorm:"column:ip;size:16" json:"ip"`                             //ip
	Addr     string `gorm:"column:addr;size:64" json:"addr"`                         //地址
	RoleID   uint   `gorm:"column:roleId" json:"roleId"`                             //用户对应角色
}
