/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 20:17:45
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 20:51:58
 * @FilePath: /gvd_server/flags/enter.go
 */
package flags

import "flag"

type Options struct {
	DB   bool   // 初始化数据库
	Port int    // 换端口
	Load string //导入数据库文件
}

func Parse() (option *Options) {
	option = new(Options)
	flag.BoolVar(&option.DB, "db", false, "初始化数据库")
	flag.IntVar(&option.Port, "port", 0, "程序运行的端口")
	flag.StringVar(&option.Load, "load", "", "导入sql数据库")
	flag.Parse()
	return option

}

// Run 根据里面的参数运行不同的脚本
func (option Options) Run() bool {
	if option.DB {
		DB()
		return true
	}

	if option.Port != 0 {
		Port(option.Port)
		return true
	}

	if option.Load != "" {
		Load()
		return true
	}
	return false
}
