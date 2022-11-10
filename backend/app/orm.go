package app

//
//import (
//	"log"
//	// "mysql-master/mysql-master"
//
//	// "github.com/go-sql-driver/mysql"
//	"gorm.io/driver/mysql"
//
//	"study.com/model"
//
//	"gorm.io/gorm"
//)
//
//type Users struct {
//	id     uint `gorm:"primaryKey"`
//	name   string
//	passwd string
//}
//
//type Todos struct {
//	id      uint `gorm:"primaryKey"`
//	userId  uint
//	title   string
//	content string
//}
//
//// 添加新的todo
//func Add(title string, content string, userid uint) (uint, error) {
//	var newTodo = model.Todos{content: content, Title: title, Userid: userid}
//	err := db.Create(&newTodo).error
//	if err != nil {
//		return 0, err
//	} else {
//		return newTodo.ID, nil
//	}
//}
//
//// 查询todo
//func get(id uint) (*model.Todo, error) {
//	var result *model.Todo
//	err := db.Where("user_id =?", id).Find(&result).Error
//	if err != nil {
//		return nil, err
//	} else {
//		return &result, nil
//	}
//}
//
//func orm() {
//	//connect mysql
//	dsn := "user:@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//create two tables
//	db.AutoMigrate(&Users{})
//	db.AutoMigrate(&Todos{})
//}
