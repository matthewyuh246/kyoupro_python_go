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

	var N int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d", &N)

	P := make([]int, N+1)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := 1; i <= N; i++ {
		P[i], _ = strconv.Atoi(line[i-1])
	}

	dp := make([]int, N+1)
	dp[1] = 0
	dp[2] = abs(P[1] - P[2])
	for i := 3; i <= N; i++ {
		dp[i] = min(abs(P[i-1]-P[i])+dp[i-1], abs(P[i-2]-P[i])+dp[i-2])
	}

	ans := []int{}
	place := N
	for {
		ans = append(ans, place)
		if place == 1 {
			break
		}

		if dp[place]-abs(P[place-1]-P[place]) == dp[place-1] {
			place -= 1
		} else {
			place -= 2
		}
	}

	ans = reverse(ans)
	fmt.Println(len(ans))

	strs := make([]string, len(ans))
	for i, v := range ans {
		strs[i] = strconv.Itoa(v)
	}
	fmt.Println(strings.Join(strs, " "))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func reverse(s []int) []int {
	result := make([]int, len(s))
	for i := range s {
		result[i] = s[len(s)-1-i]
	}

	return result
}
