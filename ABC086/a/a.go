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
	fields := strings.Fields(line)

	a, _ := strconv.Atoi(fields[0])
	b, _ := strconv.Atoi(fields[1])

	if a*b%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
