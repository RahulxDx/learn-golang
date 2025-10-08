package main

import "fmt"

func main() {
	fmt.Println("=== Multiple Return Values in Go ===")

	sum, product := calculate(5, 3)
	fmt.Println("Sum =", sum)
	fmt.Println("Product =", product)

	quotient, remainder := divide(17, 4)
	fmt.Println("Quotient =", quotient)
	fmt.Println("Remainder =", remainder)

	name, age := getPerson()
	fmt.Println("Name:", name, ", Age:", age)
}

func calculate(a int, b int) (int, int) {
	return a + b, a * b
}

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func getPerson() (string, int) {
	return "Rahul", 21
}
