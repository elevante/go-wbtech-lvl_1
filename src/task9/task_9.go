package main

import (
	"fmt"
	"time"
)

func main() {
	array := []int{2, 4, 8}
	fmt.Println(conveyor(array))
}

func conveyor(array []int) []int {

	numbersChannel := make(chan int, len(array))
	resultsChannel := make(chan int, len(array))

	go func() {
		for i := 0; i < len(array); i++ {
			numbersChannel <- array[i]
			time.Sleep(time.Second)
		}
		close(numbersChannel)
	}()

	go worker(numbersChannel, resultsChannel)

	var result []int
	for i := range resultsChannel {
		result = append(result, i)
	}
	return result
}

func worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second)
		results <- j * 2
	}
	close(results)
}
