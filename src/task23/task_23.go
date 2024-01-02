package main

import "fmt"

func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("Original slice:", numbers)

	numbers = remove(numbers, 2)
	fmt.Println("Slice after removing element:", numbers)
}
