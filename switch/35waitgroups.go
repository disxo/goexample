package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(id int, wg *sync.WaitGroup)  {
	fmt.Printf("worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
	wg.Done()  // 通知waitgroup 当前协程的工作已经完成
}

func main() {
	// 被用于等待该函数开启的所有协程
	var wg sync.WaitGroup
	// 开启几个协程，并为其递增waitgroup的计数器
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker2(i, &wg)
	}
	// 阻塞，知道waitgroup 计数器恢复为0， 即所有协程的工作都已经完成
	wg.Wait()
}
