package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

// https://www.cnblogs.com/wolfred7464/p/4670864.html

func main() {
	log.Println("jdjdj")
	c, err := redis.Dial("tcp", "192.168.204.222:6379")
	if err != nil {
		log.Printf("链接redis 失败：, Failed code: %#v", err)
	}
	log.Println(c.Do("dbsize"))
	log.Println(redis.String(c.Do("info")))
	log.Println("-----------------------------------------")
	log.Println(c.Do("GET", "guild_cross_info:1878"))
	log.Println(redis.Strings(c.Do("config", "get", "*")))
	slowlogs, _ := redis.SlowLogs(c.Do("slowlog", "get"))
	log.Println(slowlogs)

	// c.Do("lpush", "redlist", "qqq")
	// c.Do("lpush", "redlist", "www")
	// c.Do("lpush", "redlist", "eee")
	values, _ := redis.Values(c.Do("lrange", "redlist", "0", "100"))

	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}

}
