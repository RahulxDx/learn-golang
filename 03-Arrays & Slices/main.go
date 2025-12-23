package main

import "fmt"

func main() {
	// Arrays
	// var ages [3]int = [3]int{21, 27, 31}
	var ages = [3]int{21, 27, 31}

	names := [4]string{"Test1", "Test2", "Test3"}
	names[1] = "test2"
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices
	var scores = []int{100, 50, 70}
	scores[2] = 25
	scores = append(scores, 80)

	fmt.Println(scores, len(scores))

	// Slice ranges
	range1 := names[1:3]
	range2 := names[2:]
	range3 := names[:3]

	fmt.Println(range1, range2, range3)

	range1 = append(range1, "Test4")
	fmt.Println(range1)
}
