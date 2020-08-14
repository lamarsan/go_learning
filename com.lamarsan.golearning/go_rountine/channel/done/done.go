package main

import (
	"fmt"
	"sync"
)

/**
 * description channel demo
 *
 * @author lamar
 * @date 2020/7/25 5:38 下午
 */
func doWorker(id int, c chan int, w worker) {
	for n := range c {
		fmt.Printf("Worker %d Recived %d\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w.in, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
