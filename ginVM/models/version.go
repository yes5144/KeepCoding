package models

import "log"

type Version struct {
	ID        int    `json:"id,omitempty"`
	Game      string `json:"game,omitempty"`
	Zone      string `json:"zone,omitempty"`
	Version   string `json:"version,omitempty" gorm:"default:'1.0.21.22'"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func GetAllVersion() []Version {
	var mods []Version
	// Db.Select("Game").Find(&mods)
	Db.Find(&mods)
	log.Println(mods)
	return mods
}

func GetOneVersion(n int) Version {
	var one Version
	Db.First(&one, n)
	log.Println(one)
	return one
}

func NewOneVersion(v Version) bool {
	Db.Create(&v)
	return true
}
