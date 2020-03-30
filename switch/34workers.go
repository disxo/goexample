package main

import (
	"fmt"
	"time"
)

func workers(id int, jobs <- chan int, result chan <- int)  {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 这里是启动的了worker，但是没有发送任务
	for w := 1; w <= 3; w++ {
		go workers(w, jobs, results)
	}

	// 这里是发送任务
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 9; a++ {
		<- results
	}

	// 为啥会死锁呢？
	fmt.Println("result", <-results) // fatal error: all goroutines are asleep - deadlock!

}
