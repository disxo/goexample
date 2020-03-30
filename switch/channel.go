package main

import "fmt"

func main() {
	// 使用make(chan val-type) 创建一个新的通道
	messages := make(chan string)
	go func() {messages <- "ping"}()

	msg := <- messages
	fmt.Println(msg)
}
