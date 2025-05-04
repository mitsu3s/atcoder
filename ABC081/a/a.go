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

	var count int

	for _, c := range s {
		if c == '1' {
			count++
		}
	}

	fmt.Println(count)
}
