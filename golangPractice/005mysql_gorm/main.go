package main

import (
	// "encoding/json"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo ...
type UserInfo struct {
	// ID       int       `json:"id,omitempty" gorm:"primary_key"`
	// gorm.Model
	Name     string    `json:"name,omitempty"`
	Gender   string    `json:"gender,omitempty"`
	Hobby    string    `json:"hobby,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty" gorm:"type:timestamp"`
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

	db.LogMode(true)
	// 自动迁移
	db.AutoMigrate(&UserInfo{})
	u1 := UserInfo{
		Name:   "qimi2",
		Gender: "manhh",
		Hobby:  "篮球,羽毛球"}
	// u2 := UserInfo{24, "wkk", "man", "boxxing,table"}

	db.Create(&u1)
	// db.Create(&u2)

	// 查询
	var u = new(UserInfo)
	db.First(u)
	p := new(UserInfo)
	// err = json.Unmar([]byte(u), p)
	// if err != nil {
	// 	log.Fatalf("瞎jbj, Failed code %#v", err)
	// }
	log.Printf("Type: %T, value: %#v", u, u)
	log.Printf("Type: %T, value: %#v", p, p)

	// all
	var u4 = new(UserInfo)
	db.Find(&u4)
	log.Printf("Type: %T, value: %#v", u4, u4)

}
