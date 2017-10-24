package common

import (
	"errors"
	"fmt"
	"os"
	"plugin"

	"gopl.io/plugins/greeter/common/contracts"
)

var (
	// GreeterDescriptor factory
	GreeterDescriptor = map[string]contracts.IGreeter{
		"chinese": newGreeter("chinese"),
		"english": newGreeter("english"),
	}
)

// NewGreeter constructor
func NewGreeter(lang string) (contracts.IGreeter, error) {
	var (
		greeter contracts.IGreeter
		found   bool
	)

	if greeter, found = GreeterDescriptor[lang]; !found {
		return nil, errors.New("This language is not definded")
	}

	return greeter, nil
}

func newGreeter(lang string) contracts.IGreeter {
	l := "english"
	if lang != "" {
		l = lang
	}
	var mod string
	switch l {
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
	return constructor.(func(string) contracts.IGreeter)(lang)
}
