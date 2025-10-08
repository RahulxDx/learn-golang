package main

import "fmt"

func main() {
	fmt.Println("=== Booleans and Conditionals in Go ===")

	isStudent := true
	hasID := false

	fmt.Println("\n--- Boolean Values ---")
	fmt.Println("isStudent:", isStudent)
	fmt.Println("hasID:", hasID)
	fmt.Println("isStudent AND hasID:", isStudent && hasID)
	fmt.Println("isStudent OR hasID:", isStudent || hasID)
	fmt.Println("NOT isStudent:", !isStudent)

	age := 20

	fmt.Println("\n--- If-Else Statements ---")
	if age < 18 {
		fmt.Println("Minor")
	} else if age >= 18 && age < 60 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior")
	}

	score := 85

	fmt.Println("\n--- Switch Statement ---")
	switch {
	case score >= 90:
		fmt.Println("Grade A")
	case score >= 75:
		fmt.Println("Grade B")
	case score >= 50:
		fmt.Println("Grade C")
	default:
		fmt.Println("Grade D")
	}

	day := "Tuesday"

	fmt.Println("\n--- Switch with string ---")
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the week")
	default:
		fmt.Println("Midweek")
	}
}
