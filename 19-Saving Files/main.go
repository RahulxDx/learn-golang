package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Enter file name: ")
	name, _ := r.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter text to save: ")
	text, _ := r.ReadString('\n')
	text = strings.TrimSpace(text)

	file, err := os.Create(name + ".txt")
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer file.Close()

	file.WriteString(text)

	fmt.Println("File saved")
}
