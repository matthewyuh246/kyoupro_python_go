package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func linearSearch(numbers []int, value int) bool {
	for _, v := range numbers {
		if v == value {
			return true
		}
	}
	return false
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
	X, _ := strconv.Atoi(parts[1])

	sc.Scan()
	numStr := strings.Fields(sc.Text())
	A := make([]int, N+1)
	for i, s := range numStr {
		A[i], _ = strconv.Atoi(s)
	}

	if linearSearch(A, X) {
		fmt.Fprintln(wr, "Yes")
	} else {
		fmt.Fprintln(wr, "No")
	}
}
