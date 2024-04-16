package main

import (
	"fmt"

	"github.com/H-ADJI/go-data-structures/queue"
)

func main() {
	q := queue.NewQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q)
	q.Dequeue()
	q.Enqueue(3)
	fmt.Println(q)
	q.Dequeue()
	q.Enqueue(4)
	q.Enqueue(5)

	fmt.Println(q)
}
