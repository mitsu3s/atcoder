package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	t := scanner.Text()
	scanner.Scan()
	u := scanner.Text()

	tlen := len(t)
	ulen := len(u)

	found := false

	for i := 0; i <= tlen-ulen; i++ {
		ok := true
		for j := 0; j < ulen; j++ {
			if t[i+j] != '?' && t[i+j] != u[j] {
				ok = false
				break
			}
		}
		if ok {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
