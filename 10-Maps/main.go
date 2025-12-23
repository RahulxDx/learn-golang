package main

import "fmt"

func main() {
	scorecard := map[string]float64{
		"student1": 95.6,
		"student2": 98.7,
		"student3": 90.5,
	}
	fmt.Println(scorecard)
	fmt.Println(scorecard["student3"])

	// Looping Maps
	for k, v := range scorecard {
		fmt.Println(k, "-->", v)
	}
	phonebook := map[int]string{
		12345: "test_user1",
		67890: "test_user2",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[67890])
	phonebook[12345] = "testuser1"
	fmt.Println(phonebook)
}
