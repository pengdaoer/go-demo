package main

import (
	"fmt"
	"web-demo-interface-test/config"
	"web-demo-interface-test/dao/userDao"
	"web-demo-interface-test/log"
	"web-demo-interface-test/router"
)

func main() {
	s := `{
			"server": {
			  "port": 8080
			},
			"mysql": {
			  "host": "localhost",
			  "port": 3306,
			  "db": "scott",
			  "username": "root",
			  "password": "123456"
			},
			"log":{
			  "level": "debug",
			  "filename": "log/gin_blog.log",
			  "maxsize": 500,
			  "max_age": 7,
			  "max_backups": 10
			}
		  }`
	//初始化Config的全局变量
	if err := config.InitFromStr(s); err != nil {
		fmt.Printf("config.Init failed, err:%v\n", err)
		return
	}

	// 初始化日志模块
	if err := log.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	// 初始化Mysql数据库
	if err := userDao.InitMysql(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	// 初始化redis数据库
	//if err := dao.InitRedis(config.Conf.RedisConfig); err != nil {
	//	fmt.Printf("init redis failed, err:%v\n", err)
	//	return
	//}

	// 初始化
	log.Logger.Info("start project...")

	r := router.SetupRouter() // 初始化路由
	r.Run()
}
