package main

import (
	"fmt"
	"os"
	"plugin"

	"gopl.io/plugins/greeter/common/contracts"
)

func main() {
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
		fmt.Printf("Don't speak that language\n")
		os.Exit(1)
	}

	plugin, err := plugin.Open(mod)
	if err != nil {
		fmt.Printf("Cann't open a plugin %v\n", err)
		os.Exit(1)
	}

	constructor, err := plugin.Lookup("NewGreeter")
	if err != nil {
		fmt.Printf("Cann't look up a constructor %v\n", err)
		os.Exit(1)
	}

	//var greeter contracts.Greeter
	greeter := constructor.(func() contracts.Greeter)()

	greeter.Greet()

	fmt.Printf("It's working as you see\n")

	greeter.Bye()
}
