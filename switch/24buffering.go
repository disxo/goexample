package main

import "fmt"

func main() {
	// 这里我们创建了一个字符串通道，最多允许缓存2个值
	messages := make(chan string, 2)

	// 由于此通道是缓冲的，因此我们可以将这些值发送到同道中
	// 而且不需要响应的并发接收
	messages <- "buffered"
	messages <- "channel"

	// 然后我们可以像前面一样接收这两个值
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}