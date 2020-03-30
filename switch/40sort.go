package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("ints: ", ints)

	// 可以判断一个序列是不是已经排好序
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted: ", s)
}
