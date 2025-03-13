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

	var N, S int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &S)

	A := make([]int, N)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := 0; i < N; i++ {
		A[i], _ = strconv.Atoi(line[i])
	}

	dp := make([][]bool, N+1)
	for i := range dp {
		dp[i] = make([]bool, S+1)
	}

	dp[0][0] = true
	for i := 1; i < S+1; i++ {
		dp[0][i] = false
	}

	for i := 1; i < N+1; i++ {
		for j := 0; j < S+1; j++ {
			if j < A[i-1] {
				if dp[i-1][j] == true {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			}

			if j >= A[i-1] {
				if dp[i-1][j] == true || dp[i-1][j-A[i-1]] == true {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			}
		}
	}

	if dp[N][S] == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
