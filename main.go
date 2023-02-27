package main

import (
	"InstantMessaging/router"
	"InstantMessaging/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySql()
	r := router.Router()
	r.Run(":8081") //listen and server on 0.0.0.0:8080(for windows "localhost:8080")
}
