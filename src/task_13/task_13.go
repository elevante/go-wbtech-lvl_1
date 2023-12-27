package main

import "fmt"

func main() {
	fmt.Println(createSet([]string{"cat", "cat", "dog", "cat", "tree"}))

}

func createSet(array []string) []string {
	result := make(map[string]bool)

	for _, value := range array {

		result[value] = true

	}
	keys := make([]string, 0, len(result))

	for key := range result {
		keys = append(keys, key)
	}
	return keys
}
