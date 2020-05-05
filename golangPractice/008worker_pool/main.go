package main

// https://gobyexample-cn.github.io/worker-pools
import (
	"log"
	"time"
)

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 0; i < 4; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 3; i <= numJobs; i++ {
		<-results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		log.Println("worker", id, "start job", j)
		time.Sleep(time.Second)
		log.Println("worker", id, "finished job", j)
		results <- j * 2

	}
}
