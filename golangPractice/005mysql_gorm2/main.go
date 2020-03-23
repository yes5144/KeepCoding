package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Deploy struct {
	ID uint `gorm:"primary_key" json:"id,omitempty"`
	// CreatedAt   time.Time `json:"created_at,omitempty"  gorm:"default:'2006-01-02 13:04:05'"`
	// FinishedAt  time.Time `json:"finished_at,omitempty" gorm:"default:'2012-01-02 13:14:15'"`
	// DeletedAt   time.Time `json:"deleted_at,omitempty"  gorm:"default:'2099-01-06 20:20:20'"`
	ProjectName string `json:"project_name,omitempty"`
	AppName     string `json:"app_name,omitempty"`
	Version     string `json:"version,omitempty" gorm:"default:'0.2.22'"`
	Status      bool   `json:"status,omitempty" gorm:"default:1"`
	Details     string `json:"details,omitempty"`
	IsDeleted   bool   `json:"is_deleted,omitempty" gorm:"default:0"`
}

func main() {
	log.Println("hello mysql gorm")
	db, err := gorm.Open("mysql", "root:channel@tcp(192.168.204.222:3306)/gincmdb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("创建数据库链接失败, Failed code: %#v", err)
		panic("创建MySQL数据库链接失败")
	}
	defer db.Close()

	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "demo_" + defaultTableName
	// }

	db.LogMode(true)
	db.AutoMigrate(&Deploy{})

	// 设置随机种子
	// rand.Seed(time.Now().Unix())
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(100)) // [0,100)的随机值，返回值为int

	proName := []string{"天龙八部", "碧血剑", "射雕英雄传", "连城诀", "鹿鼎记", "coder"}
	appName := []string{"golang", "python", "shell", "ruby", "perl", "乔峰", "段誉", "干将", "莫邪", "黄蓉", "周伯通", "郭靖", "落花流水", "韦小宝", "陈近南"}
	log.Println(proName[rand.Intn(len(proName))])
	log.Println(appName[rand.Intn(len(appName))])

	// // var mod Deploy
	for i := 0; i < 3; i++ {
		log.Println(i, []bool{true, false, true, false}[rand.Intn(4)])
		db.Create(&Deploy{
			// FinishedAt:  time.Now(),
			// DeletedAt:   time.Now().AddDate(99, 10, 20),
			ProjectName: proName[rand.Intn(len(proName))],
			AppName:     appName[rand.Intn(len(appName))],
			Version:     "",
			Status:      []bool{true, false, true, false}[rand.Intn(4)],
			Details:     "",
			IsDeleted:   []bool{true, false, true, false}[rand.Intn(4)],
		})
	}

	log.Println("输出第一条数据。。。")
	d1 := new(Deploy)
	log.Printf("Type: %T, value: %#v", d1, d1)
	erre := db.First(d1).GetErrors()
	if erre != nil {
		log.Printf("jjd, Failed code: %#v", erre)
	}
	log.Printf("Type: %T, value: %#v", d1, d1)

	// log.Println("所有数据。。。")
	// var d2 = make([]Deploy, 0)
	// log.Printf("Type: %T, value: %#v", d2, d2)
	// db.Find(&d2)
	// log.Printf("Type: %T, value: %#v", d2, d2)

	// log.Println(db.Find(&Deploy{}))
}
