package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(x int, N int, K int, A []int) bool {
	sum := 0
	for i := 0; i < N; i++ {
		sum += x / A[i]
	}
	return sum >= K
}

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
}

func main() {
	defer wr.Flush()

	sc.Scan()
	parts := strings.Fields(sc.Text())
	N, _ := strconv.Atoi(parts[0])
	K, _ := strconv.Atoi(parts[1])

	sc.Scan()
	numStr := strings.Fields(sc.Text())
	A := make([]int, N+1)
	for i, s := range numStr {
		A[i], _ = strconv.Atoi(s)
	}

	Left := 1
	Right := int(1e9)
	for Left < Right {
		Mid := (Left + Right) / 2
		Answer := check(Mid, N, K, A)

		if !Answer {
			Left = Mid + 1
		} else {
			Right = Mid
		}
	}
	fmt.Println(Left)
}
