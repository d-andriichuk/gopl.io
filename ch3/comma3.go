package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, s := range os.Args[1:] {
		var n float64
		var err error
		if n, err = strconv.ParseFloat(s, 64); err != nil {
			fmt.Fprintf(os.Stdout, "This argument [%s] isn't an float. Try again later\n", s)
			continue
		}
		fmt.Fprintf(os.Stdout, "This number [%g] is separated by comma like this: %s\n", n, comma(s))
	}
}

func comma(value string) string {
	var buf bytes.Buffer
	cnt := (len(value)) % 3
	if cnt != 0 {
		fmt.Fprintf(&buf, "%s%s", value[:cnt], ",")
	}
	for i := cnt; i < len(value)-1; {
		fmt.Fprintf(&buf, "%s", value[i:i+3])
		if i += 3; i < (len(value) - 1) {
			fmt.Fprintf(&buf, "%s", ",")
		}
	}
	return buf.String()
}
