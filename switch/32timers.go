package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(time.Second * 2)
	<- time1.C
	fmt.Println("timer 1 expired")

	// 自己没有理解这里的作用是什么？
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}
}
