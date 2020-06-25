package main

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	Cfg      *ini.File
	RunMode  string
	RunMode2 string

	JwtSecret string
	PageSize  int
)

func main() {
	var err error
	Cfg, err = ini.Load("../configs/app.ini")
	if err != nil {
		log.Printf("load 'configs/app.ini' error, Failed code: %#v", err)
	}
	log.Println(Cfg)
	// no section
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	log.Println(RunMode)

	RunMode2 = Cfg.Section("").Key("RUN_MODE2").MustString("debug")
	log.Println(RunMode2)
	// get section app
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Printf("Fail to get section 'app', Failed code: %#v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("DFJKLEXAOO@#!")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)

	log.Println(JwtSecret)
	log.Println(PageSize)
}
