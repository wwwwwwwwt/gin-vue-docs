/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 21:02:09
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 21:51:12
 * @FilePath: /gvd_server/models/enter.go
 */
package models

import "time"

type Model struct {
	ID        uint      `json:"id" gorm:"primaryKey"` // 添加主键id
	CreatedAt time.Time `json:"createAt"`             // 添加时间
	UpdatedAt time.Time `json:"updateAt"`             // 更新时间
}
