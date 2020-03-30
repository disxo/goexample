package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"bobo", 20})

	fmt.Println(person{
		name: "Alice",
		age:  18,
	})

	fmt.Println(person{
		name: "Sam",
	})

	fmt.Println(&person{
		name: "Josh",
		age:  21,
	})

	s := person{
		name: "Bean",
		age:  50,
	}

	fmt.Println(s.name)

	sp := &s
	sp.age = 56

	fmt.Println(sp.age)
}
