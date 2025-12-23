package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Choose (a-add, t-tip, s-save, q-quit): ")
		opt, _ := r.ReadString('\n')
		opt = strings.TrimSpace(opt)

		switch opt {
		case "a":
			fmt.Print("Item name: ")
			name, _ := r.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Item price: ")
			price, _ := r.ReadString('\n')
			price = strings.TrimSpace(price)

			fmt.Println(name, price)

		case "t":
			fmt.Print("Tip amount: ")
			tip, _ := r.ReadString('\n')
			tip = strings.TrimSpace(tip)

			fmt.Println(tip)

		case "s":
			fmt.Println("Saved")

		case "q":
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}
