package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	array := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println((sampleUnion(array, 10)))
}

func sampleUnion(array []float64, step int) map[int][]float64 {

	sort.Float64s(array)
	groups := make(map[int][]float64)

	for _, value := range array {
		key := int(math.Floor(value/float64(step))) * step
		groups[key] = append(groups[key], value)
	}
	return groups
}
