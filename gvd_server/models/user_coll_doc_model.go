/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 09:50:09
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 12:47:30
 * @FilePath: /gvd_server/models/user_coll_doc_model.go
 */
package models

type UserCollDocModel struct {
	Model
	DocID     uint      `gorm:"column:docID" json:"docID"`
	DocMolel  DocModel  `gorm:"foreignKey:DocID"`
	UserID    uint      `gorm:"column:usrID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
}
