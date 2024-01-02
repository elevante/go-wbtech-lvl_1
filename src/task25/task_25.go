package main

import (
	"fmt"
	"time"
)

func sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Start")
	sleep(2)
	fmt.Println("End")
}
