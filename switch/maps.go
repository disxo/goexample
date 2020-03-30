package main

import "fmt"

// map 就是go内置关联数据类型
func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// 第二个值用来判断这个键是否存在。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}