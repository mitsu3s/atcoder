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
	q, _ := strconv.Atoi(fields[2])

	view := make([]map[int]struct{}, n+1)
	isAll := make([]bool, n+1)

	for i := 1; i <= n; i++ {
		view[i] = make(map[int]struct{})
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < q; i++ {
		scanner.Scan()
		line := scanner.Text()
		fields := strings.Fields(line)
		t, _ := strconv.Atoi(fields[0])

		if t == 1 {
			x, _ := strconv.Atoi(fields[1])
			y, _ := strconv.Atoi(fields[2])
			if !isAll[x] {
				view[x][y] = struct{}{}
			}
		} else if t == 2 {
			x, _ := strconv.Atoi(fields[1])
			isAll[x] = true
			view[x] = nil
		} else if t == 3 {
			x, _ := strconv.Atoi(fields[1])
			y, _ := strconv.Atoi(fields[2])
			if isAll[x] {
				fmt.Fprintln(writer, "Yes")
			} else {
				if _, ok := view[x][y]; ok {
					fmt.Fprintln(writer, "Yes")
				} else {
					fmt.Fprintln(writer, "No")
				}
			}
		}
	}
}
