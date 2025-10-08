package main

import "fmt"

func main() {
	fmt.Println("=== Pointers in Go ===")

	num := 10
	fmt.Println("Before function call, num =", num)

	double(&num)
	fmt.Println("After function call, num =", num)

	str := "Hello"
	fmt.Println("\nBefore function call, str =", str)

	changeString(&str)
	fmt.Println("After function call, str =", str)
}

func double(n *int) {
	*n = *n * 2
	fmt.Println("Inside function, n =", *n)
}

func changeString(s *string) {
	*s = *s + " Go"
	fmt.Println("Inside function, s =", *s)
}
