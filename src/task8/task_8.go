package main

import "fmt"

func setBit(n int64, index uint) int64 {
	return n | (1 << index)
}

func clearBit(n int64, index uint) int64 {
	return n &^ (1 << index)
}

func main() {
	var n int64 = 10
	fmt.Printf("Исходное число: %b\n", n)

	n = setBit(n, 1)
	fmt.Printf("После установки бита: %b\n", n)

	n = clearBit(n, 2)
	fmt.Printf("После сброса бита: %b\n", n)
}
