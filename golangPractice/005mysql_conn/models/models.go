package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB sqlx
var DB *sqlx.DB

// initDB ...
func init() {
	str := "root:channel@tcp(192.168.204.222:3306)/gopher?charset=utf8&parseTime=true"
	log.Println(str)
	var err error
	// DB, err = sqlx.Connect("mysql", str)
	DB, err = sqlx.Open("mysql", str)
	if err != nil {
		log.Printf("连接数据库错误, Failed code: %#v", err)
	}
	log.Println("Mysql 数据库连接成功")
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
}
