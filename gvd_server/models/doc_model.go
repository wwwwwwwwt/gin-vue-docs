/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 09:34:14
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 09:47:38
 * @FilePath: /gvd_server/models/doc_model.go
 */
package models

type DocMolel struct {
	Model
	Title       string    `gorm:"comment:文档标题" json:"title"`
	Content     string    `gorm:"comment:文档内容" json:"-"`
	DiggCount   int       `gorm:"comment:点赞量" json:"diggcount"`
	LookCount   int       `gorm:"comment:收藏量" json:"lookCount"`
	Key         string    `gorm:"comment:key;not null;unique" json:"key"`
	ParenrId    *uint     `gorm:"comment:父文档id;column:parentId" json:"parentId"`
	ParentDoc   *DocMolel `gorm:"foreignKey;comment:父文档;column:parentDoc" json:"-"` //父文档
	FreeContent string    `gorm:"comment:预览部分;column:freeContent" json:"freeContent"`
}
