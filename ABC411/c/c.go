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

	var n, q int
	fmt.Fscan(reader, &n, &q)

	a := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &a[i])
	}

	black := make([]bool, n)
	segments := 0

	for _, x := range a {
		p := x - 1
		left := p-1 >= 0 && black[p-1]
		right := p+1 < n && black[p+1]

		if !black[p] {
			switch {
			case !left && !right:
				segments++
			case left && right:
				segments--
			}
			black[p] = true
		} else {
			switch {
			case !left && !right:
				segments--
			case left && right:
				segments++
			}
			black[p] = false
		}

		fmt.Fprintln(writer, segments)
	}
}
