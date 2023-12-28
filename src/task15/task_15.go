package main

var justString string = "HELLO"

func someFunc() {
	a := 1000
	justString = a[:100]
}

func main() {
	someFunc()
}
