package main

import (
	"fmt"
	"time"
)

func main() {
	// 将所有想限制的请求的处理，我们将这些请求发送给一个相同的通道
	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)

	// 我们将每200秒加收一个值，这个是速率限制任务中的管理器
	limiter := time.Tick(time.Millisecond * 200)

	for req := range request {
		<- limiter
		fmt.Println("request", req, time.Now())
	}

	// 有时候我们想临时进行速率限制，并且不影响整体的速率控制，
	// burstyLimiter通道用来进行三次临时的脉冲型速率限制
	burstyLimiter := make(chan time.Time, 3)
	// 想将通道填充需要临时改变3次的值，要做好准备
	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每200ms我们将添加一个新的值到burstylimiter中，直到达到3个的限制
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 现在模拟超过5个接入请求。它们中刚开始的3个将
	// 由于受burstyLimiter的脉冲影响

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)
	for req := range burstyRequests {
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
