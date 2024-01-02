package main

import (
	"fmt"
	"strings"
)

func areCharsUnique(s string) bool {
	charMap := make(map[rune]bool)
	for _, char := range strings.ToLower(s) {
		if _, exists := charMap[char]; exists {
			return false
		}
		charMap[char] = true
	}
	return true
}

func main() {
	fmt.Println(areCharsUnique("abcd"))      // true
	fmt.Println(areCharsUnique("abCdefAaf")) // false
	fmt.Println(areCharsUnique("aabcd"))     // false
}
