package main

import "gopl.io/plugins/greeter/common/contracts"

// NewGreeter constructor
func NewGreeter(name string) contracts.IGreeter {
	g := &greeting{name: name}
	return g
}
