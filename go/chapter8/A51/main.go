package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	Q, _ := strconv.Atoi(sc.Text())

	var stack []string

	for i := 0; i < Q; i++ {
		sc.Scan()
		line := sc.Text()
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "1":
			stack = append(stack, parts[1])
		case "2":
			if len(stack) > 0 {
				fmt.Println(stack[len(stack)-1])
			}
		case "3":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}
}
