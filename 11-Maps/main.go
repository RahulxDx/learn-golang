package main

import "fmt"

func main() {
	fmt.Println("=== Maps in Go ===")

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	fmt.Println("Colors Map:", colors)
	fmt.Println("Red color code:", colors["red"])

	colors["yellow"] = "#FFFF00"
	fmt.Println("After adding yellow:", colors)

	fmt.Println("\n--- Iterating over Map ---")
	for key, value := range colors {
		fmt.Println(key, ":", value)
	}

	fmt.Println("\n--- Deleting from Map ---")
	delete(colors, "green")
	fmt.Println("After deleting green:", colors)

	code, exists := colors["green"]
	fmt.Println("Does green exist?", exists, "Value:", code)
}
