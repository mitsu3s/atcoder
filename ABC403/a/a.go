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
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	line := scanner.Text()
	strs := strings.Split(line, " ")

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(strs[i])
	}

	var sum int

	for i, v := range a {
		if i%2 == 0 {
			sum += v
		}
	}

	fmt.Println(sum)
}
