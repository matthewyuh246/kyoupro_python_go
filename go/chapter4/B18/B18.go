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

	if dp[N][S] == false {
		fmt.Println(-1)
		os.Exit(0)
	}

	ans := []int{}
	place := S
	for i := N; i > 0; i-- {
		if dp[i-1][place] == true {
			place = place - 0
		} else {
			place = place - A[i-1]
			ans = append(ans, i)
		}
	}
	reverse(&ans)
	ans2 := make([]string, len(ans))
	for i := 0; i < len(ans); i++ {
		ans2 = append(ans2, strconv.Itoa(ans[i]))
	}
	fmt.Println(len(ans))
	fmt.Print(strings.Join(ans2, " "))
}

func reverse(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n/2; i++ {
		(*arr)[i], (*arr)[n-1-i] = (*arr)[n-1-i], (*arr)[i]
	}
}
