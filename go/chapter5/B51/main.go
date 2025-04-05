package main

import (
	"fmt"
)

func main() {
	var kakko string

	fmt.Scan(&kakko)

	stack := []int{}

	for i, ch := range kakko {
		if ch == '(' {
			stack = append(stack, i+1)
		} else if ch == ')' {
			if len(stack) > 0 {
				openIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fmt.Println(openIndex, i+1)
			}
		}
	}
}
