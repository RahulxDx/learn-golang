package main

import "fmt"

func main() {
	b := newBill("Rahul Bill")

	b.add("soup", 4.50)
	b.add("pie", 8.95)
	b.add("pudding", 4.95)
	b.add("coffee", 3.25)

	b.setTip(10)

	fmt.Println(b.show())
}
