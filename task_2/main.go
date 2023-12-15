package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var array = [5]int{2, 4, 6, 8, 10}
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
		fmt.Println(result)
	}

}

func square(x int) int {
	time.Sleep(1 * time.Second)
	return x * x
}
