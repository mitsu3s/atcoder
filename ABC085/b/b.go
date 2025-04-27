package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	d := make([]int, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		d[i], _ = strconv.Atoi(scanner.Text())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(d)))

	uniqueD := make([]int, 0, n)

	for i, v := range d {
		if i == 0 || d[i] != d[i-1] {
			uniqueD = append(uniqueD, v)
		}
	}

	fmt.Println(len(uniqueD))
}
