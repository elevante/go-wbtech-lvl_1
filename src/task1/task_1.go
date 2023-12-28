package main

import "fmt"

type Human struct {
	Name    string
	Age     int
	IsRun   bool
	IsSpeak bool
}

type Action struct {
	Human
}

func (h *Human) Run() {
	h.IsRun = true
}

func (h *Human) Speak() {
	h.IsSpeak = true
}

func CheckHuman(h *Human) string {
	h = &Human{
		Name:    "Vladimir",
		Age:     25,
		IsRun:   false,
		IsSpeak: false,
	}
	h.Run()
	h.Speak()
	return fmt.Sprintf("Name: %s, age: %d\nRunning: %v\nSpeaking: %v\n",
		h.Name,
		h.Age,
		h.IsRun,
		h.IsSpeak)
}

func CheckAction(a *Action) string {
	a = &Action{
		Human{
			Name:    "Oleg",
			Age:     30,
			IsRun:   false,
			IsSpeak: false,
		},
	}
	a.Run()
	a.Speak()
	return fmt.Sprintf("Name: %s, age: %d\nRunning: %v\nSpeaking: %v\n",
		a.Name,
		a.Age,
		a.IsRun,
		a.IsSpeak)
}

func main() {
	var human Human
	var action Action
	fmt.Println(CheckHuman(&human))
	fmt.Println(CheckAction(&action))
}
