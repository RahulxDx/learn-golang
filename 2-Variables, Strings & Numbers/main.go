package main

import "fmt"

func main() {

	// Variables
	var name string = "Rahul"
	var age int = 21
	height := 5.9

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)

	// Numbers
	var a, b int = 10, 3
	fmt.Println("\n--- Number Operations ---")
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Remainder:", a%b)

	// Strings
	firstName := "Rahul"
	lastName := "R N"
	fullName := firstName + " " + lastName

	fmt.Println("Full Name:", fullName)
	fmt.Println("Length of Name:", len(fullName))
}
