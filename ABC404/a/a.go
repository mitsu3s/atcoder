package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	line, _ := reader.ReadString('\n')
	s = strings.TrimSpace(line)

	seen := make(map[rune]bool)
	for _, c := range s {
		seen[c] = true
	}

	for c := 'a'; c <= 'z'; c++ {
		if !seen[c] {
			fmt.Println(string(c))
			return
		}
	}
}
