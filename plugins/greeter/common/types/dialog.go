package types

import (
	"fmt"

	"gopl.io/plugins/greeter/common/contracts"
)

// Dialog type
type Dialog struct {
	dialog []string
}

// NewDialog constructor
func NewDialog() contracts.IDialog {
	return &Dialog{dialog: []string{}}
}

// Add method
func (rcv *Dialog) Add(p string) {
	rcv.dialog = append(rcv.dialog, p)
}

// Show method
func (rcv *Dialog) Show(greeter string) {
	for _, p := range rcv.dialog {
		fmt.Printf("%s: %s\n", greeter, p)
	}
}
