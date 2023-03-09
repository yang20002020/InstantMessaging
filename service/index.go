package service

import (
	"github.com/gin-gonic/gin"
	"text/template"
)

// GetIndex
// @Tags         首页
// @Success      200  {string}  welcome
// @Router      /index [get]
func GetIndex(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "welcome!!",
	//})
	ind, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
}
func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "register")
}
