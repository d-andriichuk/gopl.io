package contracts

// IGreeter interface
type IGreeter interface {
	Greet()
	Bye()
	Say(string)
	Speak(IDialog)
	GetName() string
	SetName(string)
}

// IDialog interface
type IDialog interface {
	Show(string)
	Add(string)
}
