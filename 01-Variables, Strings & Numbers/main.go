package main

import "fmt"

func main() {
	// strings
	var a string = "Test1"
	var b = "Test2"
	var c string

	fmt.Println(a, b, c)

	a = "test1"
	c = "test3"

	fmt.Println(a, b, c)

	d := "Test4"

	fmt.Println(d)

	// ints
	var age1 int = 19
	var age2 = 23
	age3 := 31

	fmt.Println(age1, age2, age3)

	// bits and memory
	var num1 int8 = 127
	var num2 int8 = -128
	var num3 uint = 25

	fmt.Println(num1, num2, num3)

	var score1 float32 = 25.98
	var score2 float64 = 8881234.66
	score3 := 1.5

	fmt.Println(score1, score2, score3)
}
