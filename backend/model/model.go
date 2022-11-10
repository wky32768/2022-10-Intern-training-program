package model

type Users struct {
	Id     uint `gorm:"primaryKey"`
	Name   string
	Passwd string
}

type Todos struct {
	Id      uint `gorm:"primaryKey"`
	UserId  uint
	Title   string
	Content string
}
