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

	line, _ = reader.ReadString('\n')
	fields := strings.Fields(line)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(fields[i])
	}

	count := 0
	for {
		isFinish := true
		for i := 0; i < n; i++ {
			if a[i]%2 != 0 {
				isFinish = false
				break
			}
			a[i] /= 2
		}
		if !isFinish {
			break
		}
		count++
	}
	fmt.Println(count)
}
