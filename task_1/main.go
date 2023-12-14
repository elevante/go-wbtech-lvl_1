package main

import "fmt"

type Human struct {
	Name  string
	Age   int
	IsRun bool
}

type Action struct {
	Human
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
	a.Run()
	fmt.Println(a.IsRun)
}
