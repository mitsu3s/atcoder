// ACx33 REx7 Incorrect answer
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
	scanner.Buffer(make([]byte, 0, 1<<20), 1<<20)

	scanner.Scan()
	line := scanner.Text()
	fields := strings.Fields(line)

	n, _ := strconv.Atoi(fields[0])
	d, _ := strconv.Atoi(fields[1])

	scanner.Scan()
	line = scanner.Text()
	strs := strings.Split(line, " ")

	a := make([]int, n)
	maxA := 0
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(strs[i])
		if a[i] > maxA {
			maxA = a[i]
		}
	}

	if d == 0 {
		seen := make(map[int]bool, n)
		distinct := 0
		for _, v := range a {
			if !seen[v] {
				seen[v] = true
				distinct++
			}
		}
		fmt.Println(n - distinct)
		return
	}

	freq := make([]int, maxA+1)
	for _, v := range a {
		freq[v]++
	}

	totalKeep := 0
	for r := 0; r < d; r++ {
		prevSkip, prevKeep := 0, 0
		for v := r; v <= maxA; v += d {
			keep := prevSkip + freq[v]
			skip := prevKeep
			if skip > keep {
				keep = skip
			}
			prevSkip, prevKeep = prevKeep, keep
		}
		totalKeep += prevKeep
	}

	fmt.Println(n - totalKeep)
}
