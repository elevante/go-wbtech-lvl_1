package main

import "fmt"

func main() {
	fmt.Println(intersectionOfManya([]int{1, 3, 7, 11}, []int{3, 11, 17, 19}))
}

func intersectionOfManya(firstArray []int, secondArray []int) []int {
	var results []int
	for _, i := range firstArray {
		for _, j := range secondArray {
			if i == j {
				results = append(results, i)
			}
		}
	}
	return results
}
