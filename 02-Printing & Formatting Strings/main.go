package main

import "fmt"

func main() {
	name := "Rahul"
	age := 21

	// Print
	fmt.Print("hello")
	fmt.Print("world!\n")

	// Println
	fmt.Println("Hello")
	fmt.Println("Bye")
	fmt.Println("My name is", name, "and my age is", age)

	// Printf (String formatting)
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %q and my name is %q\n", age, name)
	fmt.Printf("Age is of type %T\n", age)
	fmt.Printf("You weight %f\n", 58.67)
	fmt.Printf("You weight %0.1f\n", 58.67)

	// Sprintf (Saved String Formatting)
	var str = fmt.Sprintf("My age is %v \n", age)
	fmt.Println("The saved String is: ", str)
}
