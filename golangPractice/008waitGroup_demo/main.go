package main

// https://gobyexample-cn.github.io/waitgroups
import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 6; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("worker starting: ", id)
	time.Sleep(time.Second)
	log.Println("worker done.", id)

}
