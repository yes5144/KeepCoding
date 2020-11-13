package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	log.Println("hello world")
	for _, v := range []byte("wo shi") {
		log.Println(string(v))
	}

	ss := "1.23.333.23"

	news := strings.Replace(ss, "23", "255", 1)
	log.Println(news)

	// fmt.Printf
	num := 2.332
	var insertStr string
	for i := 0; i < 4; i++ {
		insertStr = insertStr + fmt.Sprintf(`("%v","%v","%v","%v",now(),now()),`, "侠客行", "xkx2", num, i)
	}
	fmt.Println(insertStr)
	insertStr = strings.TrimRight(insertStr, ",")
	fmt.Println(insertStr)
}
