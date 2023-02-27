package main

import "github.com/gin-gonic/gin"

/*
测试：http://127.0.0.1:8080/ping
*/
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() //listen and server on 0.0.0.0:8080(for windows "localhost:8080")
}
