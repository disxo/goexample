package main

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}

type rect1 struct {
	width, height float64
}

type circle1 struct {
	radius float64
}

func (r rect1) area() float64 {
	return r.width * r.height
}

func (r rect1) perim() float64 {
	return 2 * r.width + 2 * r.height
}

func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	r := rect1{
		width:  3,
		height: 4,
	}

	measure(r)
}