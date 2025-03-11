package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line1, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(line1))

	line2, _ := reader.ReadString('\n')
	A := strings.Fields(line2)

	line3, _ := reader.ReadString('\n')
	B := strings.Fields(line3)

	// DPの状態を保持する変数
	prev2, prev1 := 0, 0

	// 初期値設定
	if N >= 2 {
		prev1, _ = strconv.Atoi(A[0]) // A[2] に相当
	}

	// DP計算
	for i := 3; i <= N; i++ {
		a, _ := strconv.Atoi(A[i-2]) // A[i]
		b, _ := strconv.Atoi(B[i-3]) // B[i]
		current := min(prev1+a, prev2+b)
		prev2, prev1 = prev1, current
	}

	// 結果出力
	fmt.Println(prev1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
