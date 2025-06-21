// ACx32, TLE15, MLE2
package main

import (
	"bufio"
	"fmt"
	"os"
)

type query struct {
	t int
	p int
	s string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	pc := make([]string, n+1)
	server := ""

	for i := 0; i < q; i++ {
		var t, p int
		var s string

		fmt.Fscan(reader, &t, &p)
		if t == 1 {
			pc[p] = server
		} else if t == 2 {
			fmt.Fscan(reader, &s)
			pc[p] += s
		} else if t == 3 {
			server = pc[p]
		}
	}

	fmt.Fprintln(writer, server)
}
