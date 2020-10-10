package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Db *gorm.DB

type User struct {
	Name      string
	Telephone string
}

func InitDb() *gorm.DB {
	var err error
	Db, err = gorm.Open("sqlite3", "sqlite3.db")
	if err != nil {
		log.Printf("gorm.Open db err, Failed code %#v", err)
		panic(fmt.Sprintf("gorm.Open db err, Failed code %v", err))
	}

	Db.SingularTable(true)
	Db.LogMode(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.AutoMigrate(&User{})
	return Db
}

func (u *User) Create(user User) {
	Db.Create(&user)
}

func (u *User) Update(user User) {
	Db.Where("id=?", user.Telephone).Update(&user)
}

func (u *User) SelectByTel(tel string) {
	log.Println("before u: ", u)
	Db.Where("telephone=?", tel).First(u)
	log.Println("after u: ", u)
}

func main() {
	log.Println("hello")
	db := InitDb()
	defer db.Close()
	var user User
	user.Telephone = "11123123"
	user.Name = "xkhh"
	user.Create(user)
	fmt.Println("create")

	user.SelectByTel("123123")
	fmt.Println(user.Name)
	user.SelectByTel("34343434")
	fmt.Println(user.Name)
}
