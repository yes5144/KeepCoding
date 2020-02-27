package main

import "log"

func main() {
	log.Println("hello ")
	for i := 0; i < 3; i++ {
		log.Println("value", i)
	}
}
