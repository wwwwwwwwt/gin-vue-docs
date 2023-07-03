/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 14:27:20
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 15:09:46
 * @FilePath: /gvd_server/core/init_config.go
 */
package core

import (
	"gvd_server/config"
	"gvd_server/global"
	"os"

	"gopkg.in/yaml.v3"
)

const yamlPath = "settings.yaml"

func InitConfig() (c *config.Config) {
	byteData, err := os.ReadFile(yamlPath)
	if err != nil {
		global.Log.Fatalln("read yaml err: ", err.Error())
	}
	c = new(config.Config)
	err = yaml.Unmarshal(byteData, c)

	if err != nil {
		global.Log.Fatalln("解析 yaml err: ", err.Error())
	}
	return c
}
