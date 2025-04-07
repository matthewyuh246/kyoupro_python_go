package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binarySearch(numbers []int, value int) int {
	var binarySearchRec func(left int, right int) int
	binarySearchRec = func(left, right int) int {
		if left > right {
			return -1
		}

		mid := (left + right) / 2
		if numbers[mid] == value {
			return mid
		} else if numbers[mid] < value {
			return binarySearchRec(mid+1, right)
		} else {
			return binarySearchRec(left, mid-1)
		}
	}
	return binarySearchRec(0, len(numbers)-1)
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

	answer := binarySearch(A, X)
	wr.WriteString(fmt.Sprintf("%d\n", answer+1))
}
