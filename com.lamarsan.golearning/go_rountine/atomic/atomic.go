package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * description 锁并发
 *
 * @author lamar
 * @date 2020/8/8 4:28 下午
 */
type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Print(a.get())
}
