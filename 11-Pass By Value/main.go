package main

import "fmt"

func updateUserName(name string) string {
	name = "Rahul"
	return name
}

func updatePriceList(prices map[string]float64) {
	prices["tea"] = 1.49
}

func main() {

	customerName := "Test_user1"

	customerName = updateUserName(customerName)
	fmt.Println("Updated customer name:", customerName)

	priceList := map[string]float64{
		"sandwich": 3.25,
		"cake":     4.50,
	}

	updatePriceList(priceList)
	fmt.Println("Updated price list:", priceList)
}
