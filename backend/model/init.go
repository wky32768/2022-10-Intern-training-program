package model

import (
	"backend/app/response"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var DB *gorm.DB

// Add 添加新的todo
func Add(c echo.Context) error {
	id, error1 := strconv.Atoi(c.QueryParam("Id"))
	uid := uint(id)
	if error1 != nil {
		return response.SendResponse(c, http.StatusOK, "create failed", "error")
	}
	userid, error2 := strconv.Atoi(c.QueryParam("UserId"))
	uintUserid := uint(userid)
	if error2 != nil {
		return response.SendResponse(c, http.StatusOK, "create failed", "error")
	}
	content := c.QueryParam("Content")
	title := c.QueryParam("Title")
	var newTodo = Todos{Id: uid, UserId: uintUserid, Content: content, Title: title}
	err := DB.Create(&newTodo).Error
	if err != nil {
		return response.SendResponse(c, http.StatusOK, "create failed", "error")
	} else {
		return response.SendResponse(c, http.StatusOK, "create succeeded", "newTodo.id")
	}
}

// Get 查询todo
func Get(c echo.Context) error {
	id, error1 := strconv.Atoi(c.QueryParam("Id"))
	uid := uint(id)
	if error1 != nil {
		return response.SendResponse(c, http.StatusOK, "get failed", "error")
	}
	var result Todos
	err := DB.Where("user_id =?", uid).Find(&result).Error
	if err != nil {
		return response.SendResponse(c, http.StatusOK, "get failed", "error")
	} else {
		return response.SendResponse(c, http.StatusOK, "Todos", "get succeeded")
	}
}

// Delete 删除todo
func Delete(c echo.Context) error {
	id, error1 := strconv.Atoi(c.QueryParam("Id"))
	uid := uint(id)
	if error1 != nil {
		return response.SendResponse(c, http.StatusOK, "delete failed", "error")
	}
	var result Todos
	err := DB.Where("user_id =?", uid).Find(&result).Error
	if err != nil {
		return response.SendResponse(c, http.StatusOK, "delete failed", "error")
	} else {
		DB.Where("user_id =?", uid).Delete(result)
		return response.SendResponse(c, http.StatusOK, "result", "delete succeeded")
	}
}

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
