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
	fmt.Println("你好宇宙")
}

// Bye method
func (g *greeting) Bye() {
	fmt.Println("你好宇宙 (i don't know how to say bye in chinese. Sorry...)")
}

// Say method
func (g *greeting) Say(seq string) {
	fmt.Printf("%s greeter: %s\n", g.name, seq)
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
