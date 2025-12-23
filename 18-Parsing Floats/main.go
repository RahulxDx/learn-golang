package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Enter price: ")
	t, _ := r.ReadString('\n')
	t = strings.TrimSpace(t)

	p, err := strconv.ParseFloat(t, 64)
	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	fmt.Println("Price:", p)
}
