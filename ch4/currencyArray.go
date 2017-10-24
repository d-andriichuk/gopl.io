package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RUR
)

func main() {
	currSymb := [...]string{USD: "$", EUR: "€", GBP: "£", RUR: "₽"}
	fmt.Println(USD, currSymb[USD])
}
