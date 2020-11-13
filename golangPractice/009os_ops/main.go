package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	log.Println("hello")
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error, Failed code: %#v", err)
		return
	}
	log.Println("当前目录：", pwd)
	// cmd := exec.Command("cmd", "/C", "../scripts/iptest.bat")
	cmd := exec.Command("python ../scripts/zipByApp.py")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Start()
	log.Println(stderr.String())
	log.Println(stdout.Bytes())
	// out, err := exec.Command("python ../scripts/zipByApp.py").Output()
	if err != nil {
		log.Printf("命令执行失败：, Failed code: %#v", err)
	}
	log.Println(time.Now().Format("20060102"))
}
