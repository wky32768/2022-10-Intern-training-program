package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init 连接数据库中把Users与Todos放入数据表
func Init() {
	connectDatabase()
	fmt.Println("Database connected successfully")
	if !DB.Migrator().HasTable(&Users{}) {
		err := DB.AutoMigrate(&Users{})
		if err != nil {
			logrus.Error("cannot create user")
		}
	}
	if !DB.Migrator().HasTable(&Todos{}) {
		err := DB.AutoMigrate(&Todos{})
		if err != nil {
			logrus.Error("cannot create todo")
		}
	}
}

// 尝试连接数据库
func connectDatabase() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./model")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	loginInfo := viper.GetStringMapString("Users")

	dbArgs := loginInfo["username"] + ":" + loginInfo["password"] +
		"@(localhost)/" + loginInfo["db_name"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	if err != nil {
		logrus.Error("connect connect MySQL")
	}
}
