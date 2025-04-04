package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	kakko := sc.Text()

	var stack []int

	for i := range kakko {
		if kakko[i] == '(' {
			stack = append(stack, i+1)
		} else if kakko[i] == ')' {
			fmt.Println(kakko[0:], i+1)
		}
		kakko = kakko[1:]
	}
}
