package main

import (
	"bufio"
	"os"
	"strconv"
)

func mergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	center := len(numbers) / 2
	left := mergeSort(numbers[:center])
	right := mergeSort(numbers[center:])

	merged := make([]int, 0, len(numbers))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}

func binarySearchLeft(numbers []int, value int) int {
	left, right := 0, len(numbers)-1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if numbers[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
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
	sc.Split(bufio.ScanWords)
}

func main() {
	defer wr.Flush()

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	A := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		A[i] = a
	}

	numbers := mergeSort(A)

	sc.Scan()
	Q, _ := strconv.Atoi(sc.Text())
	for i := 0; i < Q; i++ {
		sc.Scan()
		X, _ := strconv.Atoi(sc.Text())

		ans := binarySearchLeft(numbers, X)
		wr.WriteString(strconv.Itoa(ans) + "\n")
	}
}
