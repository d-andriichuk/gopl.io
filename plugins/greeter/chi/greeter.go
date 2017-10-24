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
	fmt.Println("你好宇宙")
}

// Bye method
func (g greeting) Bye() {
	fmt.Println("你好宇宙 (i don't know how to say bye in chinese. Sorry...)")
}

// Greeter variable
//var Greeter greeting
