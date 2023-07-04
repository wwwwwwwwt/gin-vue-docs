/*
 * @Author: zzzzztw
 * @Date: 2023-07-04 20:55:49
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-04 21:09:49
 * @FilePath: /gvd_server/config/config_jwt.go
 */
package config

type Jwt struct {
	Expires int    `yaml:"expires"`
	Issuer  string `yaml:"issuer"`
	Secret  string `yaml:"secret"`
}
