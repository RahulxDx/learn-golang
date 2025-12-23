package main

import "fmt"

func main() {
	score := 50
	fmt.Println(score <= 50)
	fmt.Println(score >= 50)
	fmt.Println(score == 50)
	fmt.Println(score != 50)

	if score < 50 {
		fmt.Println("Failed in Exam")
	} else if score < 70 {
		fmt.Println("Average Mark")
	} else {
		fmt.Println("Good score")
	}
	names := []string{"test1", "test2", "test3"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at index", index)
			continue
		}
		if index > 2 {
			fmt.Println("Breaking at index", index)
			break
		}
		fmt.Println(index, value)
	}
}
