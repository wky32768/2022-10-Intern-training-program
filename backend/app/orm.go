package orm

import (
	"log"
	// "mysql-master/mysql-master"

	// "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"

	"backend/model"

	"gorm.io/gorm"
)

type Users struct {
	id     uint `gorm:"primaryKey"`
	name   string
	passwd string
}

type Todos struct {
	id      uint `gorm:"primaryKey"`
	user_id uint
	title   string
	content string
}

// 添加新的todo
func Add(title string, content string, userid uint) (uint, error) {
	var new_todo = model.Todo{content: content, Title: title, Userid: userid}
	err := db.Create(&new_todo).error
	if err != nil {
		return 0, err
	} else {
		return new_todo.ID, nil
	}
}

// 查询todo
func get(id uint) (*model.Todo, error) {
	var result *model.Todo
	err := db.Where("user_id =?", id).Find(&result).Error
	if err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}

func orm() {
	//connect mysql
	dsn := "user:@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//create two tables
	db.AutoMigrate(&Users{})
	db.AutoMigrate(&Todos{})
}
