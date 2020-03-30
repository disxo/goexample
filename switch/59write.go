package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error)  {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/Users/zvan/wzy.txt", d1, 0644)
	check(err)

	f, err := os.Create("/Users/zvan/wzy1.txt")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2) // wrote 5 bytes


	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3) // wrote 7 bytes

	f.Sync()  // 同步到磁盘

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4) // wrote 9 bytes


	w.Flush()  // 确保所有uhuancun的操作已写入底层写入器
}
