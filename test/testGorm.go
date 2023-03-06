package main

import (
	"InstantMessaging/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
https://gorm.io/zh_CN/docs/index.html
直接引用gorm.io/gorm 文档
*/

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	//数据库密码fendou2017;数据库名ginchat //charset=utf8mb4
	test := "root:fendou2017@tcp(127.0.0.1:3306)/ginchat?charset=utf8&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(test), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//UserBasic
	//test1(db)
	//Message
	test2(db)
	//Contact
	test3(db)
	//GroupBasic
	test4(db)
}
func test1(db *gorm.DB) {

	// 迁移 schema 纲要
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "申专"
	db.Create(user)

	fmt.Println(db.First(user, 1)) //根据整型主键查找

	// Update - 将 product 的 price 更新为 200
	db.Model(user).Update("PassWord", "1234")

}
func test2(db *gorm.DB) {

	// 迁移 schema 纲要
	db.AutoMigrate(&models.Message{})

}
func test3(db *gorm.DB) {
	db.AutoMigrate(&models.Contact{})
}
func test4(db *gorm.DB) {
	db.AutoMigrate(&models.GroupBasic{})
}
