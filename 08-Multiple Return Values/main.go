package main

import (
	"fmt"
	"strings"
)

func getNames(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	return names[0], names[1]
}

func main() {
	fn, sn := getNames("Rahul RN")

	fmt.Println("First name:", fn)
	fmt.Println("Second name:", sn)
}
