package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there"

	fmt.Println("Original String: ", greeting)

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hey"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	ages := []int{40, 20, 30, 35, 70, 60, 25}

	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 30)

	fmt.Println(index)
}
