package main

import (
	"InstantMessaging/router"
	"InstantMessaging/utils"
)

/*
		前后端分离引入swagger

	 资料地址: https://pkg.go.dev/github.com/swaggo/swag

使用方法：
1.下载地址：$ go install github.com/swaggo/swag/cmd/swag@latest
2. swag init 创建docs文件
3.go get -u github.com/swaggo/gin-swagger
4.go get -u github.com/swaggo/files
5. 运行 主程序
浏览器 127.0.0.1:8081/swagger/index.html，页面会显示swagger对应的页面logo
*/
func main() {
	//初始化配置文件以及数据库
	utils.InitConfig()
	utils.InitMySql()
	r := router.Router()
	r.Run(":8081")
}
