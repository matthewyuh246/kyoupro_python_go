package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	h := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		h[i], _ = strconv.Atoi(scanner.Text())
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = abs(h[0] - h[1])
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1]+abs(h[i-1]-h[i]), dp[i-2]+abs(h[i-2]-h[i]))
	}

	answer := []int{}
	place := n - 1
	for {
		answer = append(answer, place+1)
		if place == 0 {
			break
		}

		if dp[place-1]+abs(h[place-1]-h[place]) == dp[place] {
			place = place - 1
		} else {
			place = place - 2
		}
	}
	reverse(&answer)

	fmt.Println(len(answer))
	answer2 := make([]string, len(answer))
	for i := 0; i < len(answer); i++ {
		answer2 = append(answer2, strconv.Itoa(answer[i]))
	}
	fmt.Print(strings.Join(answer2, " "))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func reverse(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n/2; i++ {
		(*arr)[i], (*arr)[n-1-i] = (*arr)[n-1-i], (*arr)[i]
	}
}
