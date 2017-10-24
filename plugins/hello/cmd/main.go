package main

import (
	"log"
	"plugin"
)

func main() {
	p, err := plugin.Open("../plugins/hello.so")
	if err != nil {
		log.Fatal("cann't open a plugin")
	}

	v, err := p.Lookup("V")
	if err != nil {
		log.Fatal("cann't find a variable")
	}

	f, err := p.Lookup("F")
	if err != nil {
		log.Fatal("cann't find a function")
	}

	*v.(*int) = 7
	f.(func(int))(5)
}
