package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(counterStructer(50))

}

func counterStructer(n int) int {
	c := 0
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			m.Lock()
			c++
			m.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	return c
}
