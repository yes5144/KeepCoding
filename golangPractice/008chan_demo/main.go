package main

// https://www.zhangshengrong.com/p/wradZMEgXB/
import (
	"fmt"
	"log"
	"time"
)

type Demo struct {
	input         chan string
	output        chan string
	goroutine_cnt chan int
}

func NewDemo() *Demo {
	d := new(Demo)
	d.input = make(chan string, 80)
	d.output = make(chan string, 90)
	d.goroutine_cnt = make(chan int, 5)
	return d
}

func (this *Demo) Goroutine() {
	this.input <- time.Now().Format("2006-01-02 15:04:05")
	time.Sleep(time.Millisecond * 500)
	<-this.goroutine_cnt
}

func (this *Demo) Handle() {
	for t := range this.input {
		fmt.Println("datetime is: ", t, "goroutine count is :", len(this.goroutine_cnt))
		this.output <- t + "handle"
	}
}

func main() {
	log.Println("chan")

	demo := NewDemo()
	go demo.Handle()
	for i := 0; i < 100; i++ {
		demo.goroutine_cnt <- 1
		go demo.Goroutine()
	}
	close(demo.input)
}
