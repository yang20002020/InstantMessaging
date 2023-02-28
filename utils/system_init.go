package utils

import (
	"fmt"
	"gorm.io/driver/mysql"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("config app inited:", viper.Get("app"))
	fmt.Println("config mysql inited:", viper.Get("mysql"))
}

func InitMySql() {
	fmt.Println("*************")
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//user := models.UserBasic{}
	//DB.Find(&user)
	//fmt.Println("user", user) //
	//return DB
	fmt.Println("mysql inited")
}
