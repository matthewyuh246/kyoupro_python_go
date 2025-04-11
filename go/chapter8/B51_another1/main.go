package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

const (
	initBufSize = 1024 * 1024
	maxBufSize  = 1024 * 1024 * 1024
)

func init() {
	buf := make([]byte, initBufSize)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords)
}

func main() {
	defer wr.Flush()

	sc.Scan()
	S := sc.Text()

	stack := []int{}
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stack = append(stack, i+1)
		} else {
			wr.WriteString(fmt.Sprintf("%d %d\n", stack[len(stack)-1], i+1))
			stack = stack[:len(stack)-1]
		}
	}
}
