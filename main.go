package main

import (
	"InstantMessaging/router"
	"InstantMessaging/utils"
)

/*
1.运行主程序
2. 浏览器输入 http://127.0.0.1:8081/user/getUserList
浏览器会显示
{"message":[{"ID":1,"CreatedAt":"2023-02-27T11:22:22+08:00",
"UpdatedAt":"2023-02-27T11:22:22+08:00","DeletedAt":null,
"Name":"申专","PassWord":"1234","Email":"","Identity":"",
"ClientIp":"","ClientPort":"","LoginTime":"0001-01-01T00:00:00Z",
"HeartbeatTime":"0001-01-01T00:00:00Z",
"login_out_time":"0001-01-01T00:00:00Z",
"IsLoginOut":false,"DeviceInfo":""}]}
*/
func main() {
	utils.InitConfig()
	utils.InitMySql()
	r := router.Router()
	r.Run(":8081")
}
