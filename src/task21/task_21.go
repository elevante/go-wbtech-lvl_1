package main

import "fmt"

type OldCode struct{}

func (o *OldCode) DoSomething(i int) {
	fmt.Println(i * 2)
}

type NewCode struct{}

func (n *NewCode) DoSomething(s string) {
	fmt.Println(s + s)
}

type Adapter struct {
	oldCode *OldCode
}

func (a *Adapter) DoSomething(s string) {
	i := len(s)
	a.oldCode.DoSomething(i)
}
func main() {
	adapter := &Adapter{oldCode: &OldCode{}}
	adapter.DoSomething("Hello")
}
