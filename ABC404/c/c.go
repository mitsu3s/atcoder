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

	n, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])

	graph := make([][]int, n+1)

	for i := 0; i < m; i++ {
		line, _ = reader.ReadString('\n')
		fields = strings.Fields(line)
		a, _ := strconv.Atoi(fields[0])
		b, _ := strconv.Atoi(fields[1])

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	isTwo := true

	for i := 1; i <= n; i++ {
		if len(graph[i]) != 2 {
			isTwo = false
			break
		}
	}

	visit := make([]bool, n+1)

	queue := []int{1}
	visit[1] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		for _, to := range graph[v] {
			if !visit[to] {
				visit[to] = true
				queue = append(queue, to)
			}
		}
	}

	isConnect := true

	for i := 1; i <= n; i++ {
		if !visit[i] {
			isConnect = false
			break
		}
	}

	if isTwo && isConnect && m == n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
