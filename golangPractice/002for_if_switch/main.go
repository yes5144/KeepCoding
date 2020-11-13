package main

import "log"

func main() {
	log.Println("hello ")
	for i := 0; i < 3; i++ {
		log.Println("value", i)
	}

	// for k, v := range []string{"nz","xkx2","dj"} {
	for _, v := range []string{"nz", "xkx2", "dj"} {
		log.Println(v)
	}
}
