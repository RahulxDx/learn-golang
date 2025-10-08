package main

import "fmt"

func main() {
    // Variables
    name := "Rahul"
    age := 22
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)

    // Simple function call
    greet(name)

    // Slice example
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Numbers:", numbers)
    fmt.Println("Sum:", sum(numbers))

    // Map example
    colors := map[string]string{
        "red":   "#FF0000",
        "green": "#00FF00",
        "blue":  "#0000FF",
    }
    fmt.Println("Colors map:", colors)
}

// Function to greet
func greet(name string) {
    fmt.Println("Hello,", name)
}

// Function to sum numbers in a slice
func sum(nums []int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
