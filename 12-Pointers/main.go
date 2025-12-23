package main

import "fmt"

func set(n *string) {
	*n = "Rahul"
}

func main() {
	name := "Unknown User"

	fmt.Println(&name)

	p := &name
	fmt.Println(p)
	fmt.Println(*p)

	set(p)
	fmt.Println(name)
}
