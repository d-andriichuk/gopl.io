package main

import (
	"flag"
	"fmt"
	"gopl.io/ch2/lenconv"
	"gopl.io/ch2/tempconv"
	//"os"
	//"strconv"
)

var t = flag.Float64("t", 0.0, "temperature to convert")
var l = flag.Float64("l", 0.0, "length to convert")

func main() {
	flag.Parse()

	f := tempconv.Fahrenheit(*t)
	c := tempconv.Celsius(*t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

	if *l != 0.0 {
		km := lenconv.Kilometre(*l)
		mile := lenconv.Mile(*l)
		fmt.Printf("%s = %s, %s = %s\n", km, lenconv.KMToMile(km), mile, lenconv.MileToKM(mile))
	}
	// for _, arg := range os.Args[1:] {
	// 	t, err := strconv.ParseFloat(arg, 64)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	f := tempconv.Fahrenheit(t)
	// 	c := tempconv.Celsius(t)
	// 	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	// }
}
