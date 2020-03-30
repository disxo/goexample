package main

import (
	"fmt"
	"time"
)

func main() {
	// 打点器和定时器的区别就是，一个通过管道一个是通过range来实现
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()

	// 打点器和定时器一样，是可以停止的，将在1600ms后，停止这个打点器
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("ticker stopped")
}
