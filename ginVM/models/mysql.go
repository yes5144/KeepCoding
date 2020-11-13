package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// type Version struct {
// 	ID        int    `json:"id,omitempty"`
// 	Game      string `json:"game,omitempty"`
// 	Zone      string `json:"zone,omitempty"`
// 	Version   string `json:"version,omitempty" gorm:"default:'1.0.21.22'"`
// 	UpdatedAt string `json:"updated_at,omitempty"`
// }

var Db *gorm.DB

func init() {
	// "root:channel@tcp(192.168.204.222:3306)/gindemo?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", "root:channel@tcp(192.168.204.222:3306)/ginvm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("失败, Failed code: %#v", err)
		return
	}
	db.AutoMigrate(&Version{})
	db.LogMode(true)

	// create
	// db.Create(&Version{
	// 	Game:    "nz",
	// 	Zone:    "wx",
	// 	Version: "1.0.21",
	// })
	Db = db
	// initVersion()
}

// func GetAllVersions() {
// 	var mods []Version
// 	Db.Find(&mods)
// 	log.Println(mods)
// }

func initVersion() {
	gameZones := make(map[string][]string)
	gameZones["nz"] = []string{"wx", "oppo", "vivo", "shouq"}
	gameZones["dj"] = []string{"wx", "oppo"}
	gameZones["xkx2"] = []string{"wx", "hf"}

	for k, vs := range gameZones {
		log.Println(k, vs)
		for _, v := range vs {
			log.Println(v)
			Db.Create(&Version{
				Game: k,
				Zone: v,
			})
		}
	}
}
