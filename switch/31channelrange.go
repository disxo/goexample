package main

import "fmt"

func main() {
	// 通道中，设置两个值
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// 非空的通道也是可以关闭的
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
