package main

import (
	"InstantMessaging/router"
	"InstantMessaging/utils"
)

func main() {
	//初始化配置文件以及数据库
	utils.InitConfig()
	utils.InitMySql()
	utils.InRedis()
	r := router.Router()
	r.Run(":8081")
}
