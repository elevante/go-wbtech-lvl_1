package main

import (
	"fmt"
)

func main() {
	fmt.Println(reversesString("Hello world!"))
}

func reversesString(word string) string {
	str := []rune(word)
	var result []rune
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, str[i])
	}
	return string(result)
}
