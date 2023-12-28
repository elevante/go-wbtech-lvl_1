package main

import "fmt"

func main() {
	var str string = "Name"
	var number int = 10
	var valueBool bool
	ch := make(chan int, 1)
	ch <- 50

	definitioType(str)
	definitioType(number)
	definitioType(valueBool)
	definitioType(ch)

}

func definitioType(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Println("string:", v)
	case int:
		fmt.Println("int:", v)
	case bool:
		fmt.Println("bool:", v)
	case chan int:
		val, ok := <-v
		if ok {
			fmt.Println("channel int:", val)
		} else {
			fmt.Println("channel  is closed")
		}
	case chan string:
		val, ok := <-v
		if ok {
			fmt.Println("channel string:", val)
		} else {
			fmt.Println("channel  is closed")
		}
	default:
		fmt.Println("Not found type")
	}
}
