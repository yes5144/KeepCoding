package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/yes5144/KeepCoding/golangPractice/013crons/utils/crons"
)

func main() {
	logs.Info("hellll")
	crons.Start()
	defer crons.Stop()

	select {}

}
