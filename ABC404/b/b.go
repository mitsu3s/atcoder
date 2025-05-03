package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var line string
	line, _ = reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	s := make([]string, n)
	for i := 0; i < n; i++ {
		line, _ = reader.ReadString('\n')
		s[i] = strings.TrimSpace(line)
	}

	t := make([]string, n)
	for i := 0; i < n; i++ {
		line, _ = reader.ReadString('\n')
		t[i] = strings.TrimSpace(line)
	}

	ans := mismatch(s, t)
	now := s

	for i := 0; i < 4; i++ {
		m := mismatch(now, t)
		cost := i + m
		if cost < ans {
			ans = cost
		}
		now = rotate90(now)
	}

	fmt.Println(ans)
}

func rotate90(s []string) []string {
	n := len(s)
	result := make([]string, n)
	for i := 0; i < n; i++ {
		var row strings.Builder
		for j := n - 1; j >= 0; j-- {
			row.WriteByte(s[j][i])
		}
		result[i] = row.String()
	}
	return result
}

func mismatch(a, b []string) int {
	n := len(a)
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != b[i][j] {
				count++
			}
		}
	}
	return count
}
