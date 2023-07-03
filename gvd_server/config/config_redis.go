/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 18:57:14
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 18:59:08
 * @FilePath: /gvd_server/config/config_redis.go
 */
package config

import "fmt"

type Redis struct {
	IP       string `yaml:"addr"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	PoolSize int    `yaml:"poolsize"`
}

func (redis Redis) Addr() string {
	return fmt.Sprintf("%s:%d", redis.IP, redis.Port)
}
