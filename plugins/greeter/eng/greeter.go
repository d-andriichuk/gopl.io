package main

import (
	"fmt"

	"gopl.io/plugins/greeter/common/contracts"
)

// Greeting type
type greeting struct {
	name string
}

// NewGreeter constructor
func NewGreeter(name string) contracts.IGreeter {
	g := &greeting{name: name}
	return g
}

// Greet method
func (g *greeting) Greet() {
	fmt.Println("Hello world")
}

// Bye method
func (g *greeting) Bye() {
	fmt.Println("Good bye")
}

// Say method
func (g *greeting) Say(seq string) {
	fmt.Printf("%s greeter: %v\n", g.name, seq)
}

// Speak method
func (g *greeting) Speak(d contracts.IDialog) {
	d.Show(g.name)
}

// GetName method
func (g *greeting) GetName() string {
	return g.name
}

// SetName method
func (g *greeting) SetName(name string) {
	g.name = name
}

// Greeter variable
//var Greeter greeting
