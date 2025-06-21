package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var p string
	var l int

	fmt.Fscan(reader, &p, &l)

	if len(p) >= l {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
