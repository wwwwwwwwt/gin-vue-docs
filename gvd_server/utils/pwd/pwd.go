/*
 * @Author: zzzzztw
 * @Date: 2023-07-08 11:54:45
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-08 12:01:59
 * @FilePath: /gvdoc/gvd_server/utils/pwd/pwd.go
 */
package pwd

import (
	"gvd_server/global"

	"golang.org/x/crypto/bcrypt"
)

func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		global.Log.Error(err)
	}
	return string(hash)
}

func CheckPwd(hashpwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpwd), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}
