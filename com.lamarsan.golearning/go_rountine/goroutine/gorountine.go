package goroutine

import (
	"fmt"
	"time"
)

/**
 * description go_routine demo
 *
 * @author lamar
 * @date 2020/7/25 5:18 下午
 */
func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from "+"goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
