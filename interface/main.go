package main

import "fmt"

type Speaker interface {
	SayHello()
	SayBye()
}

type Rasum interface {
	Thinking()
}

type Human struct {
	Greeting string
	Rasum
}

func (h *Human) SayHello() {
	fmt.Println(h.Greeting)
}

func (h *Human) SayBye() {
	fmt.Println("Bye bye")
}

func (h *Human) Thinking() {
	fmt.Println("I'm thinking")
}

type Dog struct {
	Greeting string
	Byeing   string
}

func (d *Dog) SayHello() {
	fmt.Println(d.Greeting)
}

func (d *Dog) SayBye() {
	fmt.Println(d.Byeing)
}

func main() {
	var s Speaker
	var r Rasum

	human := &Human{
		Greeting: "Hello",
	}

	s = human
	s.SayHello()
	s.SayBye()

	r = human
	r.Thinking()

	s = &Dog{
		Greeting: "Woof-woof",
		Byeing:   "wooooooooooooooooo",
	}
	s.SayHello()
	s.SayBye()
}
