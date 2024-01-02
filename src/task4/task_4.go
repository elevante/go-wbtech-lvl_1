package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Data struct {
	WorkerCount int
}

func (d *Data) read() int {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the number of workers: ")
	input.Scan()
	value, err := strconv.Atoi(input.Text())
	if err != nil {
		fmt.Println("Invalid value entered")
	} else {
		d.WorkerCount = value
	}

	return d.WorkerCount
}

func main() {
	var data Data
	data.read()

	jobs := make(chan int)
	results := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < data.WorkerCount; i++ {
		go worker(ctx, i+1, jobs, results)
	}

	go func() {
		for i := 0; ; i++ {
			jobs <- i + 1
			time.Sleep(time.Second)
		}
	}()

	for i := 0; ; i++ {
		fmt.Printf("result #%d : value = %d\n", i+1, <-results)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("close")
		cancel()
		os.Exit(0)
	}()

}

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int) {

	for {
		select {
		case j, ok := <-jobs:
			if !ok {
				return
			}
			time.Sleep(time.Second)
			fmt.Printf("worker #%d finished\n", id)
			results <- j
		case <-ctx.Done():
			return
		}
	}
}
