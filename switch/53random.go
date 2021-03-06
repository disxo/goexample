package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print(rand.Intn(100), ",") // 81
	fmt.Print(rand.Intn(100)) // 87
	fmt.Println()

	fmt.Println(rand.Float64()) // 0.6645600532184904

	fmt.Print((rand.Float64() * 5) + 5, ",") // 7.1885709359349015
	fmt.Print((rand.Float64() * 5) + 5) // 7.123187485356329
	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",") // 83
	fmt.Print(r1.Intn(100)) // 84
	fmt.Println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",") // 5
	fmt.Print(r2.Intn(100)) // 87
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",") // 5
	fmt.Print(r3.Intn(100)) // 87
}
