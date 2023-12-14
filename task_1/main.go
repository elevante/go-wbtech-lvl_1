package main

import "fmt"

type Human struct {
	Name  string
	Age   int
	IsRun bool
}

type Action struct {
	Human Human
}

func (h *Human) Run() {
	fmt.Println("running")
	h.IsRun = true
}

func main() {
	h := &Human{}
	h.Run()
	fmt.Println(h.IsRun)
	a := &Action{}
	a.Human.Run()
	fmt.Println(a.Human.IsRun)
}
