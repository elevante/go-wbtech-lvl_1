package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	squares := outputSquares([]int{2, 4, 6, 8, 10})
	fmt.Println(squares)
}

func outputSquares(array []int) []int {
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

	squares := make([]int, 0, len(array))
	for result := range results {
		squares = append(squares, result)
	}
	return squares
}

func square(value int) int {
	time.Sleep(1 * time.Second)
	return value * value
}
