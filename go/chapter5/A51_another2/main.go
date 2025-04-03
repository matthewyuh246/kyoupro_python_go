package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type IStack interface {
	Push(data any)
	Pop()
	Print()
}

type Stack struct {
	stack []any
}

func NewStack() IStack {
	return &Stack{stack: []any{}}
}

func (s *Stack) Push(data any) {
	s.stack = append(s.stack, data)
}

func (s *Stack) Pop() {
	if len(s.stack) == 0 {
		return
	}
	lastIndex := len(s.stack) - 1
	s.stack = s.stack[:lastIndex]
}

func (s *Stack) Print() {
	if len(s.stack) == 0 {
		return
	}
	lastIndex := len(s.stack) - 1
	fmt.Println(s.stack[lastIndex])
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	Q, _ := strconv.Atoi(sc.Text())

	stack := NewStack()

	for i := 0; i < Q; i++ {
		sc.Scan()
		queryType := sc.Text()

		if queryType == "1" {
			sc.Scan()
			title := sc.Text()
			stack.Push(title)
		} else if queryType == "2" {
			stack.Print()
		} else if queryType == "3" {
			stack.Pop()
		}
	}
}
