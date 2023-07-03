/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 14:24:39
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 19:02:46
 * @FilePath: /gvd_server/config/enter.go
 */
package config

type Config struct {
	System System `yaml:"system"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
}
