package service

import (
	"InstantMessaging/models"
	"InstantMessaging/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"

	"math/rand"

	"strconv"
	"time"
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
		"code":    0, //0成功， -1 失败
		"message": data,
		"data":    data,
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
	//user.Name = c.Query("name")
	//passWord := c.Query("password")
	//repassWord := c.Query("repassword")
	user.Name = c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("Identity")
	fmt.Println(user.Name, "  >>>>>>>>>>>  ", password, repassword)
	//salt 随机数
	salt := fmt.Sprintf("%06d", rand.Int31())
	fmt.Println("salt:", salt)
	//
	data := models.FindUserByName(user.Name)
	//注意是“”，不是“ ”
	if user.Name == "" || password == "" || repassword == "" {
		fmt.Println("data.Name:", data.Name)
		c.JSON(200, gin.H{
			"code":    -1, //0成功， -1 失败
			"message": "用户名或密码不能为空！",
			"data":    user,
		})
		return
	}
	if data.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "用户名已注册！",
			"data":    user,
		})
		return
	}
	if password != repassword {
		c.JSON(200, gin.H{
			"code":    -1, //0成功， -1 失败
			"message": "两次密码不一致",
			"data":    user,
		})
		return
	}

	//user.PassWord=passWord
	user.PassWord = utils.MakePassword(password, salt)
	fmt.Println("user.PassWord:", password)
	user.Salt = salt
	fmt.Println("密码一致")
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0, //0成功， -1 失败
		"message": "新增用户成功！",
		"data":    user,
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
		"code":    0, //0成功， -1 失败
		"message": "删除用户成功！",
		"data":    user,
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
			"code":    -1, //0成功， -1 失败
			"message": "修改用户参数不匹配！",
			"data":    user,
		})
		return
	}

	models.UpDateUser(user)
	c.JSON(200, gin.H{
		"code":    0, //0成功， -1 失败
		"message": "修改用户成功！",
		"data":    user,
	})
}

// FindUserByNameAndPwd
// Summary 所有用户
// @Tags  用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success  200  {string}  json{"code","message"}
// @Router  /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}
	//name := c.Query("name")
	//password := c.Query("password")

	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1, //0 成功 -1  失败
			"message": "该用户不存在",
			"data":    data,
		})
		return
	}
	fmt.Println("user:", user)
	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1, //0 成功 -1  失败
			"message": "密码不正确",
			"data":    data,
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)
	c.JSON(200, gin.H{
		"code":    0, //0 成功 -1  失败
		"message": "登录成功",
		"data":    data,
	})
}

// 防止跨域站点伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("upGrade.Upgrade err.....:", err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(ws, c)

}
func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {

		msg, err := utils.Subsribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println("utils.Subsribe....:", err)
		}
		fmt.Println("发送消息:", msg)
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			fmt.Println(err)
		}
	}
}
func SendUserMsg(c *gin.Context) {
	fmt.Println("调用SendUserMsg")
	models.Chat(c.Writer, c.Request)
}
