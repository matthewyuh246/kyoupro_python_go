package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	Q, _ := strconv.Atoi(scanner.Text())

	var S []string

	for i := 0; i < Q; i++ {
		scanner.Scan()
		queryType := scanner.Text()

		if queryType == "1" {
			scanner.Scan()
			title := scanner.Text()
			S = append(S, title)

		} else if queryType == "2" {
			if len(S) > 0 {
				fmt.Fprintln(writer, S[len(S)-1])
			}

		} else if queryType == "3" {
			if len(S) > 0 {
				S = S[:len(S)-1]
			}
		}
	}
}
