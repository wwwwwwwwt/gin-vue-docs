/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 18:27:44
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 18:32:31
 * @FilePath: /gvd_server/models/user_pwd_doc_model.go
 */
package models

type UserPwdDocModel struct {
	Model
	UserID uint `gorm:"column:userID" json:"userID"`
	DocID  uint `gorm:"column:docID" json:"docID"`
}
