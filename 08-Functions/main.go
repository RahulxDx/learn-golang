package main

import "fmt"

func main() {
	fmt.Println("=== Functions in Go ===")

	result := add(5, 3)
	fmt.Println("Add(5, 3) =", result)

	fmt.Println("Subtract(10, 4) =", subtract(10, 4))

	fmt.Println("\n--- Function with Multiple Returns ---")
	sum, product := sumAndProduct(4, 5)
	fmt.Println("Sum =", sum, ", Product =", product)

	fmt.Println("\n--- Function with Variadic Parameters ---")
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Total =", total)

	fmt.Println("\n--- Anonymous Function ---")
	square := func(x int) int {
		return x * x
	}
	fmt.Println("Square(6) =", square(6))
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func sumAndProduct(a int, b int) (int, int) {
	return a + b, a * b
}

func sumAll(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
