package main

import (
	"fmt"
	"os"

	"gopl.io/plugins/greeter/common"
	"gopl.io/plugins/greeter/common/contracts"
	"gopl.io/plugins/greeter/common/types"
)

func main() {
	var (
		greeter contracts.IGreeter
		// found   bool
		err error
	)
	lang := "english"
	if len(os.Args) == 2 {
		lang = os.Args[1]
	}

	// if greeter, found = common.GreeterDescriptor[lang]; !found {
	// 	fmt.Printf("%v\n", errors.New("Greeter not found"))
	// 	os.Exit(1)
	// }

	if greeter, err = common.NewGreeter(lang); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	dialog := types.NewDialog()
	dialog.Add("Hello again")
	dialog.Add("Now i can use dialogs")
	dialog.Add("That's all")

	greeter.Greet()

	greeter.SetName(lang)

	greeter.Say(fmt.Sprintf("my name is %s", greeter.GetName()))
	greeter.Say("It's working, as you see")

	greeter.Speak(dialog)

	greeter.Bye()

	if lang == "english" {
		lang = "chinese"
	} else {
		lang = "english"
	}

	if greeter, err = common.NewGreeter(lang); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	greeter.Say(fmt.Sprintf("Hi, my name is %s", greeter.GetName()))
}
