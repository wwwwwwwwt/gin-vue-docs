/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 17:29:23
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 18:45:47
 * @FilePath: /gvd_server/config/config_mysql.go
 */
package config

import "fmt"

type Mysql struct {
	Host     string `yaml:"host"`     // 服务器地址:端口
	Port     int    `yaml:"port"`     //端口
	Config   string `yaml:"config"`   // 高级配置
	DB       string `yaml:"db"`       // 数据库名
	Username string `yaml:"username"` // 数据库用户名
	Password string `yaml:"password"` // 数据库密码
	LogLevel string `yaml:"logLevel"` // 是否开启Gorm全局日志
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.DB, m.Config)
}
