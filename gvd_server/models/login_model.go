/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 18:35:24
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 19:36:25
 * @FilePath: /gvd_server/models/login_model.go
 */
package models

//用户登录数据
type LoginModel struct {
	Model
	UserID    uint      `gorm:"column:userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:20" json:"ip"` //登录ip
	NickName  string    `gorm:"column:nickname;size:42" json:"nickName"`
	UA        string    `json:"ua"` //ua
	Token     string    `gorm:"size:256" json:"token"`
	Device    string    `gorm:"size:256" json:"device"` //登陆设备
	Addr      string    `gorm:"size:64" json:"addr"`
}
