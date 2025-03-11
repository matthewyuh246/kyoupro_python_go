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

	a := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	b := make([]int, n-2)
	for i := 0; i < n-2; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = a[0]

	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1]+a[i-1], dp[i-2]+b[i-2])
	}

	answer := []int{}
	place := n
	for true {
		answer = append(answer, place)
		if place == 1 {
			break
		}

		if dp[place-2]+a[place-2] == dp[place-1] {
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

func reverse(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n/2; i++ {
		(*arr)[i], (*arr)[n-1-i] = (*arr)[n-1-i], (*arr)[i]
	}
}
