package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, s := range os.Args[1:] {
		var n int
		var err error
		if n, err = strconv.Atoi(s); err != nil {
			fmt.Fprintf(os.Stdout, "This argument [%s] isn't an integer. Try again later\n", s)
			continue
		}
		fmt.Fprintf(os.Stdout, "This number [%d] is separated by comma like this: %s\n", n, comma(s))
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
