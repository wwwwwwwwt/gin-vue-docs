/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 09:50:09
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 10:07:31
 * @FilePath: /gvd_server/models/user_coll_doc_model.go
 */
package models

type UserCollDocModel struct {
	Model
	DocId     uint      `gorm:"column:docId" json:"docId"`
	DocMolel  DocMolel  `gorm:"foreignKey;DocId" json:"docModel"`
	UserID    uint      `gorm:"column:usrID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey;UserID"`
}
