package crons

import (
	"github.com/astaxie/beego/logs"
	"github.com/robfig/cron"
)

var crontab = cron.New()

func Start() {
	logs.Info("crontab start ...")
	crontab.Start()
}

func Stop() {
	logs.Info("crontab stop ...")
	crontab.Stop()
}
