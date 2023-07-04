/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 17:13:57
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 18:20:20
 * @FilePath: /gvd_server/models/image_model.go
 */
package models

import "fmt"

type ImageModel struct {
	Model
	UserID    uint      `gorm:"column:userID;comment:用户id" json:"userID"`
	RoleModel RoleModel `gorm:"foreignKey:UserID" json:"-"`
	FileName  string    `gorm:"column:fileNmae;size:128" json:"filename"`
	Size      int64     `gorm:"column:size;comment:文件大小，单位字节" json:"size"`
	Path      string    `gorm:"column:path;size:128;comment:文件路径" json:"path"`
	Hash      string    `gorm:"column:hash;size:64;comment:文件的哈希" json:"hash"`
}

func (image ImageModel) WebPath() string {
	return fmt.Sprintf("/%s", image.Path)
}
