package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	N, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	scanner.Scan()
	aStrs := strings.Fields(scanner.Text())
	A := make([]int, len(aStrs))
	for i, s := range aStrs {
		A[i], _ = strconv.Atoi(s)
	}

	scanner.Scan()
	bStrs := strings.Fields(scanner.Text())
	B := make([]int, len(bStrs))
	for i, s := range bStrs {
		B[i], _ = strconv.Atoi(s)
	}

	dp := make([]int, N+1)
	dp[1] = 0
	if N >= 2 {
		dp[2] = A[0]
	}
	for i := 3; i <= N; i++ {
		dp[i] = min(dp[i-1]+A[i-2], dp[i-2]+B[i-3])
	}

	fmt.Println(dp[N])
}
