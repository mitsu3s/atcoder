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

	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)
	n, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])

	line, _ = reader.ReadString('\n')
	fields = strings.Fields(line)
	c := make([]int64, n)

	for i := 0; i < n; i++ {
		ci, _ := strconv.ParseInt(fields[i], 10, 64)
		c[i] = ci
	}

	species2zoos := make([][]int, m)
	for i := 0; i < m; i++ {
		line, _ = reader.ReadString('\n')
		fields = strings.Fields(line)
		k, _ := strconv.Atoi(fields[0])
		species2zoos[i] = make([]int, k)
		for j := 0; j < k; j++ {
			a, _ := strconv.Atoi(fields[1+j])
			species2zoos[i][j] = a - 1
		}
	}

	pow3 := make([]int, n+1)
	pow3[0] = 1
	for i := 1; i <= n; i++ {
		pow3[i] = pow3[i-1] * 3
	}

	const INF = int64(4e18)
	best := INF

	total := pow3[n]
	for comb := 0; comb < total; comb++ {
		cost := int64(0)

		for i := 0; i < n; i++ {
			vi := (comb / pow3[i]) % 3
			if vi > 0 {
				cost += int64(vi) * c[i]
				if cost >= best {
					break
				}
			}
		}
		if cost >= best {
			continue
		}

		ok := true
		for s := 0; s < m; s++ {
			cnt := 0
			for _, i := range species2zoos[s] {
				cnt += (comb / pow3[i]) % 3
				if cnt >= 2 {
					break
				}
			}
			if cnt < 2 {
				ok = false
				break
			}
		}
		if ok && cost < best {
			best = cost
		}
	}

	fmt.Println(best)
}
