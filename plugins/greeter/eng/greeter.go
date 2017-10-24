package main

import (
	"fmt"

	"gopl.io/plugins/greeter/common/contracts"
)

// Greeting type
type greeting string

// NewGreeter constructor
func NewGreeter() contracts.Greeter {
	var g greeting
	return &g
}

// Greet method
func (g greeting) Greet() {
	fmt.Println("Hello world")
}

// Bye method
func (g greeting) Bye() {
	fmt.Println("Good bye")
}

// Greeter variable
//var Greeter greeting
