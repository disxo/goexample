package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error)  {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("/Users/zvan/wzy.txt")
	check(err)
	fmt.Print(string(dat))
	/*
	we are the champion.
	we can win the world.

	do you believe it?
	tell me
	we can
	5 bytes
	*/

	f, err := os.Open("/Users/zvan/wzy.txt")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1)) // 5 bytes: we ar


	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2)) // 2 bytes @ 6:  t


	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3)) // 2 bytes @ 6:  t

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4)) // 5 bytes: we ar

	f.Close()
}
