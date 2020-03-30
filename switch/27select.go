package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// 各个通道将在若干时间后接受一个值，这个用来模拟例如并行go协程中阻塞的rpc操作
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- c1:
			fmt.Println("received", msg1)
		case msg2 := <- c2:
			fmt.Println("received", msg2)


		}
	}
}
