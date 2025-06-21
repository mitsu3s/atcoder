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

	var n int
	fmt.Fscan(reader, &n)

	d := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &d[i])
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sum := 0
			for k := i; k < j; k++ {
				sum += d[k]
			}
			fmt.Fprint(writer, sum)
			if j < n-1 {
				fmt.Fprint(writer, " ")
			}
		}
		fmt.Fprintln(writer)
	}
}
