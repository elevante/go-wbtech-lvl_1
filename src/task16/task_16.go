package main

import (
	"fmt"
	"sort"
)

type slice interface {
	int | float64
}

func main() {
	listInts := []int{1, 5, 3, 7, 9}
	itemInt := 9
	sort.Ints(listInts)

	listFloats := []float64{1.0, 5.0, 3.0, 7.0, 9.0}
	itemFloat := 9.0
	sort.Float64s(listFloats)

	fmt.Println(binarySearch(listInts, itemInt))
	fmt.Println(binarySearch(listFloats, itemFloat))

	// Implementation with built-in methods

	iInt := sort.SearchInts(listInts, itemInt)
	fmt.Println(iInt)

	if iInt < len(listInts) && listInts[iInt] == itemInt {
		fmt.Printf("found %d at index %d in listInt\n", itemInt, iInt)
	} else {
		fmt.Printf("%d not found in listInt\n", itemInt)
	}

	xFloat := 9.0
	iFloat := sort.SearchFloat64s(listFloats, xFloat)
	if iFloat < len(listFloats) && listFloats[iFloat] == itemFloat {
		fmt.Printf("found %.1f at index %d in listFloat\n", itemFloat, iFloat)
	} else {
		fmt.Printf("%.1f not found in listFloat\n", itemFloat)
	}
}

func binarySearch[V slice](list []V, item V) int {
	var result int
	low := 0
	hight := len(list) - 1

	for low <= hight {
		mid := (low + hight) / 2
		guess := list[mid]
		if guess == item {
			result = mid
			break
		}
		if guess > item {
			hight = mid - 1
		}
		if guess < item {
			low = mid + 1
		}
	}
	return result
}
