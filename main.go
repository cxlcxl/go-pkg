// io.Reader 接口示例
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	once := new(sync.Once)
	for i := 0; i < 10; i++ {
		once.Do(func() {
			getInstance(i)
		})
		fmt.Println("main: ", i)
	}

	time.Sleep(time.Second * 2)
}

func getInstance(i int) {
	fmt.Println("once: ", i)
}
