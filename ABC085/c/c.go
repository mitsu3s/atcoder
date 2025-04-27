package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()
	fields := strings.Fields(line)

	n, _ := strconv.Atoi(fields[0])
	y, _ := strconv.Atoi(fields[1])

	var a, b, c int = -1, -1, -1

	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			k := n - i - j
			if 10000*i+5000*j+1000*k == y {
				a, b, c = i, j, k
			}
		}
	}

	fmt.Println(a, b, c)
}
