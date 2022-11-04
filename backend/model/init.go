package model

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 在连接后的数据库放入数据表
func Init() {
	connectDatabase()
	err := DB.AutoMigrate(&User{})
	if err != nil {
		logrus.Error("cannot create user")
	}
	err = DB.AutoMigrate(&Todo{})
	if err != nil {
		logrus.Error("cannot create todo")
	}
}

// 尝试连接数据库
func connectDatabase() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	loginInfo := viper.GetStringMapString("User")

	dbArgs := loginInfo["username"] + ":" + loginInfo["password"] +
		"@(localhost)/" + loginInfo["db_name"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	if err != nil {
		logrus.Error("connect connect MySQL")
	}
}
