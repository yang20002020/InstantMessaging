package models

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLoginOut    bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
