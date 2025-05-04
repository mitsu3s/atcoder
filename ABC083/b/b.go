package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func digitSum(x int) int {
	dsum := 0
	for x > 0 {
		dsum += x % 10
		x /= 10
	}
	return dsum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	line := scanner.Text()
	fields := strings.Fields(line)

	n, _ := strconv.Atoi(fields[0])
	a, _ := strconv.Atoi(fields[1])
	b, _ := strconv.Atoi(fields[2])

	var sum int

	for i := 0; i <= n; i++ {
		s := digitSum(i)
		if a <= s && s <= b {
			sum += i
		}
	}

	fmt.Println(sum)
}
