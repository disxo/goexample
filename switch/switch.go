package main

import (
	"fmt"
	"time"
)

func main()  {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Thursday, time.Wednesday:
		fmt.Println("it is the thursday, wednesday")
	case time.Sunday:
		fmt.Println("it is sunday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it is before noon")
	default:
		fmt.Println("it is after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am a bool")
		case int:
			fmt.Println("i am a int")
		default:
			fmt.Println("do not know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
