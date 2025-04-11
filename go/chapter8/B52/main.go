package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
	N, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	X, _ := strconv.Atoi(sc.Text())
	A := strings.Split(sc.Text(), "")

	queue := []int{X - 1}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		A[pos] = "@"

		if pos-1 >= 0 && A[pos-1] == "." {
			A[pos-1] = "@"
			queue = append(queue, pos-1)
		}
		if pos+1 < N && A[pos+1] == "." {
			A[pos+1] = "@"
			queue = append(queue, pos+1)
		}
	}

	wr.WriteString(strings.Join(A, ""))
}
