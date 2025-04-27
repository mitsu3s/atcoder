package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	sort.Sort(sort.IntSlice(a))

	var alice, bob int

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			alice += a[n-i-1]
			a = a[:n-i-1]
		} else {
			bob += a[n-i-1]
			a = a[:n-i-1]
		}
	}

	fmt.Println(alice - bob)
}
