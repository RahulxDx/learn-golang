package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getIn(p string, r *bufio.Reader) string {
	fmt.Print(p)
	t, _ := r.ReadString('\n')
	return strings.TrimSpace(t)
}

func main() {
	r := bufio.NewReader(os.Stdin)

	n := getIn("Name: ", r)
	fmt.Println("Hello", n)
}
