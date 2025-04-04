package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type IQueue interface {
	Enqueue(data any)
	Dequeue()
	Print()
}

type Queue struct {
	queue []any
}

func NewQueue() IQueue {
	return &Queue{queue: []any{}}
}

func (q *Queue) Enqueue(data any) {
	q.queue = append(q.queue, data)
}

func (q *Queue) Dequeue() {
	if len(q.queue) > 0 {
		q.queue = q.queue[1:]
	}
}

func (q *Queue) Print() {
	if len(q.queue) > 0 {
		fmt.Println(q.queue[0])
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	Q, _ := strconv.Atoi(sc.Text())

	queue := NewQueue()

	for i := 0; i < Q; i++ {
		sc.Scan()
		queryType := sc.Text()

		if queryType == "1" {
			sc.Scan()
			title := sc.Text()
			queue.Enqueue(title)
		} else if queryType == "2" {
			queue.Print()
		} else if queryType == "3" {
			queue.Dequeue()
		}
	}
}
