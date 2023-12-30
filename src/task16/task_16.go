package main

import "fmt"

func main() {
	list := []int{2, 1, 4, 6, 8, 0}
	fmt.Println(quikSort(list))
}

func quikSort(list []int) []int {
	if len(list) < 2 {
		return list
	} else {
		privot := list[0]
		var less []int
		var greater []int
		for _, value := range list[1:] {
			if value <= privot {
				less = append(less, value)
			}
			if value > privot {
				greater = append(greater, value)
			}
		}
		less = quikSort(less)
		greater = quikSort(greater)
		less = append(less, privot)
		return append(less, greater...)
	}
}
