package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	strs := strings.Fields(line)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(strs[i])
	}

	line, _ = reader.ReadString('\n')
	fields := strings.Fields(line)

	m, _ := strconv.Atoi(fields[0])
	p, _ := strconv.Atoi(fields[1])
	q, _ := strconv.Atoi(fields[2])

	_ = m
	_ = p
	_ = q
}
