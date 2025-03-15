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
	sc.Buffer(make([]byte, 1000), 1000000)

	var N, W int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &W)

	w := make([]int, N+1)
	v := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		sc.Scan()
		parts := strings.Fields(sc.Text())
		w[i], _ = strconv.Atoi(parts[0])
		v[i], _ = strconv.Atoi(parts[1])
	}

	const NEG_INF int = -1e9
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
		for j := 0; j < W+1; j++ {
			dp[i][j] = NEG_INF
		}
	}
	dp[0][0] = 0

	for i := 1; i < N+1; i++ {
		for j := 0; j < W+1; j++ {
			if j < w[i] {
				dp[i][j] = dp[i-1][j]
			}
			if j >= w[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
			}
		}
	}
	result := maxArray(dp[N])
	fmt.Print(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArray(arr []int) int {
	tmp := 0
	for _, v := range arr {
		if tmp > v {
			continue
		}
		tmp = v
	}
	return tmp
}
