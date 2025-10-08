package main

import (
	"fmt"
)

var globalVar = "I am a global variable"

func main() {
	fmt.Println("=== Package Scope in Go ===")

	localVar := "I am a local variable"
	fmt.Println(globalVar)
	fmt.Println(localVar)

	showMessage()
	fmt.Println("Accessing globalVar from main:", globalVar)
}

func showMessage() {
	fmt.Println("Accessing globalVar from showMessage:", globalVar)
	// fmt.Println(localVar) // This will cause an error because localVar is not in scope
}
