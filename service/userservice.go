package service

import (
	"InstantMessaging/models"
	"InstantMessaging/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
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
	//
	salt := fmt.Sprintf("%06d", rand.Int31())
	fmt.Println("salt:", salt)
	//
	data := models.FindUserByName(user.Name)
	//注意是“”，不是“ ”
	if data.Name != "" {
		fmt.Println("data.Name:", data.Name)
		c.JSON(-1, gin.H{
			"message": "用户名已经注册",
		})
		return
	}
	if passWord != repassWord {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}

	//user.PassWord=passWord
	user.PassWord = utils.MakePassword(passWord, salt)
	fmt.Println("user.PassWord:", passWord)
	user.Salt = salt
	fmt.Println("密码一致")
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
// @param phone formData string false "phone"
// @param email  formData string false "email"
// @Success  200  {string}  json{"code","message"}
// @Router  /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	fmt.Println("update:", user)
	//校验
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "修改用户参数不匹配！",
		})
		return
	}

	models.UpDateUser(user)
	c.JSON(200, gin.H{
		"message": "修改用户成功！",
	})
}

// FindUserByNameAndPwd
// Summary 所有用户
// @Tags  用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success  200  {string}  json{"code","message"}
// @Router  /user/findUserByNameAndPwd [get]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}
	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"message": "该用户不存在",
		})
		return
	}
	fmt.Println("user:", user)
	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"message": "密码不正确",
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)
	c.JSON(200, gin.H{
		"message": data,
	})
}
