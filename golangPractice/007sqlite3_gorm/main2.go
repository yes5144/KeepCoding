package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Version xxx
type Version struct {
	ID        uint
	CreateAt  string
	Project   string
	Zone      string
	Version   string
	IsDeleted bool
}

func main() {
	log.Println("hello gorm sqlite")
	// var db *gorm.DB

	db, err := gorm.Open("sqlite3", "sqlite3.db")
	if err != nil {
		log.Fatalf("open sqlite3, Failed code: %#v", err)
	}

	db.LogMode(true)
	db.AutoMigrate(&Version{})
	log.Println("初始化data")
	InitVersion(db)

}

// InitVersion xxx
func InitVersion(db *gorm.DB) {
	for _, p := range []string{"nz", "dj"} {
		log.Println(p)
		for _, z := range []string{"wx", "qq"} {
			log.Println(z)
			db.Create(&Version{
				Project: p,
				Zone:    z,
				Version: "1.2.2.33",
			})
		}
	}
}
