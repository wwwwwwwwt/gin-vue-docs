/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 09:34:14
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 13:18:29
 * @FilePath: /gvd_server/models/doc_model.go
 */
package models

type DocModel struct {
	Model
	Title           string      `gorm:"comment:文档标题" json:"title"`
	Content         string      `gorm:"comment:文档内容" json:"-"`
	DiggCount       int         `gorm:"comment:点赞量" json:"diggcount"`
	LookCount       int         `gorm:"comment:浏览量" json:"lookCount"`
	Key             string      `gorm:"comment:key;not null;unique" json:"key"`
	ParentID        *uint       `gorm:"comment:父文档id;column:parentID" json:"parentID"`
	ParentModel     *DocModel   `gorm:"foreignKey:ParentID" json:"-"` //父文档
	Child           []*DocModel `gorm:"foreignKey:ParentID" json:"-"` //他会有子孙文档
	FreeContent     string      `gorm:"comment:预览部分;column:freeContent" json:"freeContent"`
	UserCollDocList []UserModel `gorm:"many2many:user_coll_doc_models;joinForeignKey:DocID;JoinRederences:UserID" json:"-"`
}
