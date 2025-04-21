package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	limit := int(math.Sqrt(float64(number)))
	for i := 2; i < limit+1; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
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
	Q, _ := strconv.Atoi(sc.Text())
	X := make([]int, Q)
	for i := range X {
		sc.Scan()
		X[i], _ = strconv.Atoi(sc.Text())
	}

	for _, x := range X {
		if isPrime(x) {
			fmt.Fprintln(wr, "Yes")
		} else {
			fmt.Fprintln(wr, "No")
		}
	}
}
