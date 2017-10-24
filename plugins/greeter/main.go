package main

import (
	"fmt"
	"os"

	"gopl.io/plugins/greeter/common"
	"gopl.io/plugins/greeter/common/contracts"
)

func main() {
	var (
		greeter contracts.Greeter
		found   bool
	)
	lang := "english"
	if len(os.Args) == 2 {
		lang = os.Args[1]
	}

	if greeter, found = common.GreeterDescriptor[lang]; !found {
		fmt.Printf("No such language\n")
		os.Exit(1)
	}

	greeter.Greet()

	fmt.Printf("It's working, as you see\n")

	greeter.Bye()
}
