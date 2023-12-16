package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Data struct {
	WorkerCount int
}

func (d *Data) Read() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите число воркеров: ")
	input.Scan()
	value, err := strconv.Atoi(input.Text())
	if err != nil {
		fmt.Println("Invalid value entered")
	} else {
	}
	d.WorkerCount = value
}

func main() {
	var data Data
	data.Read()

	const jobsCount = 50
	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)

	for i := 0; i < data.WorkerCount; i++ {
		go worker(i+1, jobs, results)
	}

	for i := 0; i < jobsCount; i++ {
		jobs <- i + 1
	}

	close(jobs)

	for i := 0; i < jobsCount; i++ {
		fmt.Printf("result #%d : value = %d\n", i+1, <-results)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Printf("worker #%d finished\n", id)
		results <- j
	}
}
