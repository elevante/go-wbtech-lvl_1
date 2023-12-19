package main

import (
	"fmt"
	"time"
)

func main() {
	var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)

	go func(arr [10]int) {
		for _, value := range arr {
			ch <- value
			time.Sleep(2 * time.Second)
		}
		close(ch)
	}(array)

	go func() {
		for i := 0; i < 10; i++ {
			result := <-ch
			fmt.Printf("Result: %v\n", result)
		}
	}()
	<-time.After(time.Second * 5)
}
