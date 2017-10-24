package main

import (
	"fmt"

	"gopl.io/ch6/intset"
)

func main() {
	set := intset.IntSet{}
	set.AddAll(20, 17, 25)

	fmt.Println(set.Has(20))
	fmt.Println(set.Has(5))

	fmt.Printf("Init a set %s\n", &set)

	set.Clear()

	fmt.Println("Clear a set " + set.String())

	set.AddAll(15, 63, 105)

	fmt.Printf("Extend a set %s\n", &set)
	fmt.Printf("Base set's len %d\n", set.Len())

	newS := set.Copy()

	fmt.Printf("New set %s\n", newS)

	set.Remove(63)

	fmt.Printf("Remove 63 from base set %s\n", &set)
	fmt.Printf("Base set's len %d\n", set.Len())

	elems := set.Elems()

	fmt.Printf("Decimal slice of set %v\n", elems)
}
