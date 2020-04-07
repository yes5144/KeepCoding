package main

import (
	"encoding/json"
	"log"
)

// https://www.cnblogs.com/ycyoes/p/5398796.html

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"shanghai_VPN","serverIP":"192.168.204.13"},
						{"serverName":"Beijing_VPN","serverIP":"172.16.18.20"}]}
	`
	json.Unmarshal([]byte(str), &s)
	log.Println(s)
	log.Println(s.Servers[0].ServerIP)
}
