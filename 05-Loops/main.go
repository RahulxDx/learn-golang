package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 5 {
		fmt.Println("Value of x is: ", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is: ", i)
	}
	names := []string{"test1", "test2", "test3"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Println(index, value)
	}
}
