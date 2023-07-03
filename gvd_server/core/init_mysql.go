/*
 * @Author: zzzzztw
 * @Date: 2023-07-03 17:31:14
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-03 18:33:14
 * @FilePath: /gvd_server/core/init_gorm.go
 */
package core

import (
	"fmt"
	"gvd_server/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn

func InitMysql() *gorm.DB {

	if global.Config.Mysql.Host == "" {
		global.Log.Warn("未配置mysql,取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface

	var logLevel = logger.Error

	switch global.Config.Mysql.LogLevel {
	case "info":
		logLevel = logger.Info
	case "warn":
		logLevel = logger.Warn
	case "error":
		logLevel = logger.Error
	default:
		logLevel = logger.Error
	}

	mysqlLogger = logger.Default.LogMode(logLevel)
	/*
		if global.Config.System.Env == "dev" {
			// 开发环境显示所有的sql
			mysqlLogger = logger.Default.LogMode(logger.Info)
		} else {
			mysqlLogger = logger.Default.LogMode(logger.Error) // 只打印错误的sql
		}
		//global.MysqlLog = logger.Default.LogMode(logger.Info)
	*/
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("[%s] mysql连接失败", dsn))
		global.Log.Fatalf(fmt.Sprintf("[%s] mysql连接失败, err: %s", dsn, err.Error()))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout
	return db
}
