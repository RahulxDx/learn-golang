package main

import "fmt"

func main() {
	fmt.Println("=== Arrays and Slices in Go ===")

	var arr = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", arr)
	fmt.Println("First element:", arr[0])
	fmt.Println("Length of array:", len(arr))

	fmt.Println("\n--- Slice Example ---")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", nums)
	fmt.Println("Length:", len(nums))
	fmt.Println("Capacity:", cap(nums))

	fmt.Println("\n--- Slicing an Array ---")
	slice1 := arr[1:4]
	fmt.Println("Slice1 (arr[1:4]):", slice1)

	fmt.Println("\n--- Appending to Slice ---")
	nums = append(nums, 6, 7)
	fmt.Println("After append:", nums)

	fmt.Println("\n--- Copying Slices ---")
	copySlice := make([]int, len(nums))
	copy(copySlice, nums)
	fmt.Println("Copied Slice:", copySlice)
}
