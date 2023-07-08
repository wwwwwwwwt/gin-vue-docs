/*
 * @Author: zzzzztw
 * @Date: 2023-07-08 12:02:32
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-08 12:04:08
 * @FilePath: /gvdoc/gvd_server/testdata/pwf.go
 */
package main

import (
	"fmt"
	"gvd_server/utils/pwd"
)

func main() {
	str := "ztw"
	hashstr := pwd.HashPwd(str)
	ok := pwd.CheckPwd(hashstr, str)

	fmt.Println(ok)
}
