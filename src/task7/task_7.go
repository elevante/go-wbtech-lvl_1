package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(writeToMap())
}

func writeToMap() map[int]int {
	writes := 100
	storage := make(map[int]int, writes)

	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			storage[i] = i
		}(i)
	}
	wg.Wait()
	return storage
}
