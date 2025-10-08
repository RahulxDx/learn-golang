package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== The Standard Library in Go ===")

	fmt.Println("\n--- math package ---")
	fmt.Println("Pi:", math.Pi)
	fmt.Println("Square root of 16:", math.Sqrt(16))
	fmt.Println("Power (2^3):", math.Pow(2, 3))

	fmt.Println("\n--- strings package ---")
	str := "hello go"
	fmt.Println("Original:", str)
	fmt.Println("Uppercase:", strings.ToUpper(str))
	fmt.Println("Contains 'go':", strings.Contains(str, "go"))
	fmt.Println("Replace:", strings.ReplaceAll(str, "go", "Golang"))

	fmt.Println("\n--- time package ---")
	now := time.Now()
	fmt.Println("Current Time:", now)
	fmt.Println("Year:", now.Year())
	fmt.Println("Formatted:", now.Format("02-Jan-2006 15:04:05"))
}
