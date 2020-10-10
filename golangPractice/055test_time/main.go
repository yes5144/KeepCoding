package main

import (
	"time"

	"github.com/astaxie/beego/logs"
)

func main() {
	logs.Info("hello")
	nowtime := time.Now()
	logs.Info(nowtime)
}
