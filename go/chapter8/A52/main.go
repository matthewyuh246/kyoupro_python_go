package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	sc.Scan()
	Q, _ := strconv.Atoi(sc.Text())

	var T []string

	for i := 0; i < Q; i++ {
		sc.Scan()
		queryType := sc.Text()

		if queryType == "1" {
			sc.Scan()
			title := sc.Text()
			T = append(T, title)
		} else if queryType == "2" {
			if len(T) > 0 {
				fmt.Fprintln(writer, T[0])
			}
		} else if queryType == "3" {
			if len(T) > 0 {
				T = T[1:]
			}
		}
	}
}
