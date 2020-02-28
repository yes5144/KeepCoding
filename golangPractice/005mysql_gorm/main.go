package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo ...
type UserInfo struct {
	ID     int
	Name   string
	Gender string
	Hobby  string
}

// https://www.liwenzhou.com/posts/Go/gorm/
// https://gorm.io/zh_CN/docs/models.html

// 注意 数据库编码格式，go是utf8，建议mysql使用utf8mb4

func main() {
	log.Println("hell gorm")
	db, err := gorm.Open("mysql", "root:channel@(192.168.204.222:3306)/gopher?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("创建连接mysql失败, Failed code: %#v", err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&UserInfo{})
	u1 := UserInfo{23, "qimi", "man", "篮球,羽毛球"}
	u2 := UserInfo{24, "wkk", "man", "boxxing,table"}

	db.Create(&u1)
	db.Create(&u2)

	// 查询
	var u = new(UserInfo)
	db.First(u)
	log.Printf("Type: %T, value: %#v", u, u)

	// all
	var u4 = new(UserInfo)
	db.Find(&u4)
	log.Printf("Type: %T, value: %#v", u4, u4)

}
