package commit

import (
	"fmt"
	"log"
	// "mysql-master/mysql-master"

	// "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type Users struct {
	id     uint `gorm:"primaryKey"`
	name   string
	passwd string
}

type Todos struct {
	id      uint `gorm:"primaryKey"`
	userId  uint
	title   string
	content string
}

func main() {
	//connect mysql
	dsn := "user:@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//create two tables
	db.AutoMigrate(&Users{})
	db.AutoMigrate(&Todos{})

	//create
	cd := Users{id: 1, name: "cd", passwd: "juan"}
	result := db.Create(&cd)
	fmt.Println(result)

	//search
	result = db.Where("name =?", "cd").Find(&Users{})

	//modify
	temp := Users{}
	db.Where("name = ?", "cd").Take(&temp)
	temp.name = "czy"
	db.Save(&temp)

	//delete
	db.Where("name = ?", "cd").Delete(&Users{})

}
