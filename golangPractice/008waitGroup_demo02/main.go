package main

import (
	"log"
	"sync"
	"time"
)

var (
	wg       sync.WaitGroup
	jobsChan = make(chan int, 5)
)

func main() {
	log.Println("hello")
	// 生产者
	go producer(10)

	// 消费者
	consumer()

	wg.Wait()
	// close(jobsChan)

}

func producer(n int) {
	log.Println("包子。。。")
	for i := 0; i < n; i++ {
		jobsChan <- i
		log.Println("包子库存数量---", len(jobsChan))
	}
}

func consumer() {
	for x := range jobsChan {
		wg.Add(1)
		time.Sleep(time.Second)
		log.Println("消灭包子：", x)

		if len(jobsChan) == 0 {
			log.Println("nothing...")
			close(jobsChan)
			break
		}
		// _, ok := <-jobsChan
		// if !ok {
		// 	log.Println(", Something wrong!")
		// 	panic(ok)
		// }
		wg.Done()
	}
}
