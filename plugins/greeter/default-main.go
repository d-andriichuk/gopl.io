package main

import (
	"fmt"
	"log"
	"os"
	"plugin"
)

// Greeter  interface
type Greeter interface {
	Greet()
}

func main() {
	// determinate module to load
	lang := "english"
	if len(os.Args) == 2 {
		lang = os.Args[1]
	}
	var mod string
	switch lang {
	case "english":
		mod = "./eng/eng.so"
	case "chinese":
		mod = "./chi/chi.so"
	default:
		fmt.Printf("Don't know this lang")
		os.Exit(1)
	}

	// load module
	// 1. open the .so file to load the symbols
	plugin, err := plugin.Open(mod)
	if err != nil {
		log.Fatal("cann't open a plugin")
	}

	// 2. look up a symbols (exported variable or function)
	// in this case, variable Greeter
	symGreeter, err := plugin.Lookup("Greeter")
	if err != nil {
		log.Fatal("cann't find a variable")
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		log.Fatal("unexpected type from module symbol")
	}

	// 4. use the module
	greeter.Greet()
}
