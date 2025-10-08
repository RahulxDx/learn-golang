package main

import "fmt"

func main() {

	fmt.Print("Hello")
	fmt.Print(" Go!\n")
	fmt.Println("Welcome to Go!")

	name := "Rahul"
	age := 21
	height := 5.9

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)

	fmt.Printf("My name is %s and I am %d years old.\n", name, age)
	fmt.Printf("My height is %.1f feet.\n", height)

	fmt.Printf("String: %s\n", name)
	fmt.Printf("Integer: %d\n", age)
	fmt.Printf("Float (2 decimals): %.2f\n", height)
	fmt.Printf("Boolean: %t\n", true)
	fmt.Printf("Type of variable: %T\n", name)

	message := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)

	fmt.Println(message)
}
