package main

import (
	"log"
)

func main() {
	log.Println("hello world")
	for _, v := range []byte("wo shi shui") {
		log.Println(string(v))
	}
}
