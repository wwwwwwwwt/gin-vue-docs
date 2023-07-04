/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 20:39:50
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 19:44:42
 * @FilePath: /gvd_server/flags/flags_db.go
 */
package flags

import (
	"fmt"
	"gvd_server/global"
	"gvd_server/models"
)

func DB() {
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserModel{},
		&models.RoleModel{},
		&models.DocModel{},
		&models.UserCollDocModel{},
		&models.RoleDocModel{},
		&models.ImageModel{},
		&models.UserPwdDocModel{},
		&models.LoginModel{},
		&models.DocDataModel{},
	)
	if err != nil {
		global.Log.Fatalf("数据库迁移失败,err: %s", err.Error())
	}
	fmt.Println("初始化数据库")
}
