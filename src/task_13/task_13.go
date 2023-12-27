package main

import "fmt"

func main() {
	a := 2
	b := 5
	fmt.Println(changePlaces(a, b))
}

func changePlaces(a, b int) (int, int) {
	a, b = b, a
	return a, b
}
