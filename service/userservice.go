package service

import (
	"InstantMessaging/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetUserList
// Summary 所有用户
// @Tags  用户模块
// @Success  200  {string}  json{"code","message"}
// @Router  /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
	fmt.Println("data->", data)
}

// CreateUser
// Summary 新增用户
// @Tags  用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success  200  {string}  json{"code","message"}
// @Router  /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	passWord := c.Query("password")
	repassWord := c.Query("repassword")
	if passWord != repassWord {
		c.JSON(200, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	fmt.Println("密码一致")
	user.PassWord = passWord
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "新增用户成功！",
	})

}

// DeleteUser
// Summary 删除用户
// @Tags  用户模块
// @param id query string false "id"
// @Success  200  {string}  json{"code","message"}
// @Router  /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功！",
	})

}

// UpdateUser
// Summary 修改更户
// @Tags  用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password  formData string false "password"
// @Success  200  {string}  json{"code","message"}
// @Router  /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	models.UpDateUser(user)
	c.JSON(200, gin.H{
		"message": "修改用户成功！",
	})
}
