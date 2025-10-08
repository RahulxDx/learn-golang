package main

import "fmt"

func main() {
	fmt.Println("=== Loops in Go ===")

	fmt.Println("\n--- Basic for loop ---")
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", i)
	}

	fmt.Println("\n--- for loop as while ---")
	j := 1
	for j <= 5 {
		fmt.Println("j =", j)
		j++
	}

	fmt.Println("\n--- Infinite loop with break ---")
	k := 1
	for {
		fmt.Println("k =", k)
		k++
		if k > 5 {
			break
		}
	}

	fmt.Println("\n--- Loop with continue ---")
	for m := 1; m <= 5; m++ {
		if m%2 == 0 {
			continue
		}
		fmt.Println("Odd m =", m)
	}

	fmt.Println("\n--- Nested loops ---")
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 2; y++ {
			fmt.Println("x =", x, ", y =", y)
		}
	}
}
