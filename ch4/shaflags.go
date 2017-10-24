package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var sha = flag.Int("sha", 256, "type of crypto: 256, 384, 512")

func main() {
	flag.Parse()
	d := make(map[string][]byte)
	args := flag.Args()
	switch *sha {
	case 384:
		for _, arg := range args {
			f := sha512.Sum384([]byte(arg))
			d[arg] = f[:]
		}
	case 512:
		for _, arg := range args {
			f := sha512.Sum512([]byte(arg))
			d[arg] = f[:]
		}
	default:
		for _, arg := range args {
			f := sha256.Sum256([]byte(arg))
			d[arg] = f[:]
		}
	}
	for _, arg := range args {
		fmt.Fprintf(os.Stdout, "String:%v\nSHA%d digest:%x\n", arg, *sha, d[arg])
	}
}
