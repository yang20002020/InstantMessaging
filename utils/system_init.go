package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

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
	//自定义日志模板，打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)
	fmt.Println("*************")
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
		&gorm.Config{Logger: newLogger})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//user := models.UserBasic{}
	//DB.Find(&user)
	//fmt.Println("user", user) //
	//return DB
	fmt.Println("mysql inited")
}

/*
redis:

	addr: "192.168.0.10:6379"
	password: ""
	DB: 0
	poolSize: 30
	minIdConn: 30
*/
func InRedis() {
	Red := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdConn"),
	})
	pong, err := Red.Ping().Result()
	if err != nil {
		fmt.Println("init redis err...:", err)
		return
	}
	fmt.Println("redis inited ...", pong)

}
