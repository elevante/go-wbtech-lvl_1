package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int64 = 196
	binary := strconv.FormatInt(int64(number), 2)
	fmt.Println(binary)
}

func SetBit(num int64, i uint, val bool) int64 {
	mask := int64(1) << 1
	if val {

	}
}
