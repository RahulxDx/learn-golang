package main

import (
	"fmt"
	"math"
)

func sayHello(n string) {
	fmt.Printf("Hello %s\n", n)
}
func sayHelloToAll(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}
func areaCircle(r float64) float64 {
	return math.Pi * r * r
}
func main() {
	// sayHello("Test1")
	// sayHello("Test2")
	sayHelloToAll([]string{"test_user1", "test_user2", "test_user3"}, sayHello)
	a1 := areaCircle(3)
	fmt.Printf("Area of circle is %v", a1)
}
