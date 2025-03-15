package main

import (
	"fmt"
)

func main() {
	var N, S int
	fmt.Scan(&N)
	fmt.Scan(&S)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	dp := make([]int, S+1)
	dp[0] = -1
	for i, a := range A {
		for j := S; j >= a; j-- {
			if dp[j-a] != 0 && dp[j] == 0 {
				dp[j] = i + 1
			}
		}
	}

	if dp[S] == 0 {
		fmt.Println(-1)
		return
	}

	ans := make([]int, 0, S)
	for i := S; i > 0; i -= A[dp[i]-1] {
		ans = append([]int{dp[i]}, ans...)
	}

	fmt.Println(len(ans))
	for _, a := range ans {
		fmt.Print(a, " ")
	}
}
