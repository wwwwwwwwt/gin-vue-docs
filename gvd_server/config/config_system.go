/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 14:22:13
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 14:23:22
 * @FilePath: /gvd_server/config/config_system.go
 */
package config

import "fmt"

type System struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

func (system System) Addr() string {
	return fmt.Sprintf("%s:%d", system.IP, system.Port)
}
