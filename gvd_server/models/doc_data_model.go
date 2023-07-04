/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 19:38:42
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 19:43:52
 * @FilePath: /gvd_server/models/doc_data_model.go
 */
package models

type DocDataModel struct {
	Model
	DocID     uint   `gorm:"column:docID" json:"docID"`
	DocTitle  string `gorm:"column:docTitle" json:"docTitle"`
	LookCount int    `gorm:"column:lookCount" json:"lookCount"`
	DiggCount int    `gorm:"column:diggCount" json:"diggCount"`
	CollCount int    `gorm:"column:collCount" json:"collCount"`
}
