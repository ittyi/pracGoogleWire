package main

import "fmt"

type Message string

type Greeter struct {
	// ... TBD
	Message Message // <- adding a Message field
	Messag  Message // <- adding a Message field
}

type Event struct {
	// ... TBD
	Greeter Greeter // <- adding a Greeter field
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message, l Message) Greeter {
	return Greeter{Message: m, Messag: l}
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	message := NewMessage()
	fmt.Printf("massege:%s\n", message)
	greeter := NewGreeter(message, message)
	fmt.Printf("greeter:%s:%#v:%T\n", greeter, greeter, greeter)
	// event := NewEvent(greeter)

	// event.Start()
}
