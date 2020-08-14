package main

import (
	"fmt"
	"time"
)

/**
 * description channel demo
 *
 * @author lamar
 * @date 2020/7/25 5:38 下午
 */
func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d Recived %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	channelClose()
}
