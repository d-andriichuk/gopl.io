package common

import (
	"errors"
	"fmt"
	"plugin"

	"gopl.io/plugins/greeter/common/contracts"
)

const (
	chinesePlugin = "./chi/chi.so"
	englishPlugin = "./eng/eng.so"

	englishLang = "english"
	chineseLang = "chinese"
	defaultLang = englishLang
)

var (
	// GreeterDescriptor factory
	GreeterDescriptor = map[string]contracts.IGreeter{
		englishLang: newGreeter(englishLang),
		chineseLang: newGreeter(chineseLang),
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

	if greeter == nil {
		return nil, errors.New("Create greeter failed")
	}

	return greeter, nil
}

func newGreeter(lang string) contracts.IGreeter {
	l := englishLang
	if lang != "" {
		l = lang
	}
	var mod string
	switch l {
	case englishLang:
		mod = englishPlugin
	case chineseLang:
		mod = chinesePlugin
	default:
		fmt.Printf("Don't speak that language\n")
		return nil
	}

	plugin, err := plugin.Open(mod)
	if err != nil {
		fmt.Printf("Cann't open a plugin %v\n", err)
		return nil
	}

	constructor, err := plugin.Lookup("NewGreeter")
	if err != nil {
		fmt.Printf("Cann't look up a constructor %v\n", err)
		return nil
	}

	//var greeter contracts.Greeter
	return constructor.(func(string) contracts.IGreeter)(lang)
}
