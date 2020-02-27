package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configs struct {
	Mysqlconfig interface{} `json:"mysqlconfig"`
	Redisconfig interface{} `json:"redisconfig"`
}

type Mysqlconfig struct {
	Dbhost string
	Dbname string
	Dbuser string
	Dbpwd  string
	Dbport int
}

type Redisconfig struct {
	Rhost string
	Rport int
	Rauth string
	Rdb   int
}

func main() {
	log.Println("json.fil  Unmarshal")
	contents, _ := ioutil.ReadFile("../configs/config.json")
	log.Printf("Type: %T, value: %#v", contents, string(contents))
	log.Println("------------------")

	var msg Configs
	err := json.Unmarshal(contents, &msg)
	if err != nil {
		log.Printf("Configs Unmarshal, Failed code: %#v", err)
	}
	log.Println(msg)

}

// https://blog.csdn.net/qq_35191331/article/details/79696821
// https://zhuanlan.zhihu.com/p/62195508
// https://blog.csdn.net/zxy_666/article/details/80173288
// https://studygolang.com/articles/6742
