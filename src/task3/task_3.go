package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println(outputSumSquares([]int{2, 4, 6, 8, 10}))
}

func outputSumSquares(array []int) int {
	var sum int
	var wg sync.WaitGroup
	results := make(chan int, len(array))

	for _, value := range array {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			results <- square(value)
		}(value)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		sum += result
	}
	return sum
}

func square(value int) int {
	time.Sleep(1 * time.Second)
	return value * value
}
