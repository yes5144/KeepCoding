package main

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func newPool(server, passwd string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   200,
		IdleTimeout: 200 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				log.Printf("redis 不为所动, Failed code: %#v", err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("ping")
			return err
		},
	}
}

func main() {
	pool = newPool("192.168.204.222:6379", "")
	for {
		time.Sleep(time.Second)
		conn := pool.Get()
		conn.Do("set", "abc", 300)
		r, err := redis.Int(conn.Do("get", "abc"))
		if err != nil {
			log.Printf("do get 'abc' failed, Failed code: %#v", err)
			continue
		}
		log.Printf("redis key 'abc' --> Type: %T, value: %#v", r, r)
	}
}
