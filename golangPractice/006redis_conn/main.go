package main

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/yes5144/KeepCoding/golangPractice/006redis_conn/rmodels"
)

func main() {
	log.Println("hello redis")
	//
	err := rmodels.Redisdb.Set("score", 100, 0).Err()
	if err != nil {
		log.Printf("redis set scroe err, Failed code: %#v", err)
		return
	}

	val, err := rmodels.Redisdb.Get("score").Result()
	if err != nil {
		log.Printf("redis get score err, Failed code: %#v", err)
		return
	}
	log.Println("redis get score value: ", val)

	// redisExample
	redisExampl()

	// redisExample2
	redisExample2()

}

func redisExampl() {
	val, err := rmodels.Redisdb.Get("name").Result()
	if err == redis.Nil {
		log.Println("redis key: 'name' not exsit")
	} else if err != nil {
		log.Printf("redis get name err, Failed code: %#v", err)
		return
	}
	log.Println("name's value: ", val)

}

// redisExampl2
func redisExample2() {
	languageRank := "language_rank"
	languages := []*redis.Z{
		&redis.Z{Score: 90.0, Member: "Golang"},
		&redis.Z{Score: 80.0, Member: "C/C++"},
		&redis.Z{Score: 70.0, Member: "JavaScript"},
		&redis.Z{Score: 88.0, Member: "Java"},
		&redis.Z{Score: 93.0, Member: "Python"},
	}

	// ZADD
	num, err := rmodels.Redisdb.ZAdd(languageRank, languages...).Result()
	if err != nil {
		log.Printf("zadd err, Failed code: %#v", err)
	}
	log.Println(num)

	// 把golang 的分数加10
	newScore, err := rmodels.Redisdb.ZIncrBy(languageRank, 10.0, "Golang").Result()
	if err != nil {
		log.Printf("redis zincreby failed, Failed code: %#v", err)
		return
	}
	log.Println("new score is ", newScore)

	// 取分数最高的3个
	ret, err := rmodels.Redisdb.ZRevRangeWithScores(languageRank, 0, 3).Result()
	if err != nil {
		log.Printf("zrevrange fialed , Failed code: %#v", err)
		return
	}
	for k, v := range ret {
		log.Println("分数排行：", k+1, v.Member, v.Score)
	}
}

// 详情链接：https://www.liwenzhou.com/posts/Go/go_redis/
