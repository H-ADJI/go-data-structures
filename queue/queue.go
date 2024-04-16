package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	buffer   []int
	front    int
	rear     int
	capacity int
}

func NewQueue(capacity int) *Queue {
	queue := new(Queue)
	if capacity > 0 {
		buffer := make([]int, capacity)
		queue.buffer = buffer
		queue.capacity = capacity
	} else {
		queue.capacity = -1
	}

	queue.front = -1
	queue.rear = -1

	return queue
}

func (q *Queue) Enqueue(element int) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	if q.IsEmpty() {
		q.front++
		q.rear++
		q.buffer[q.rear] = element
		return nil
	}
	q.rear++
	q.buffer[q.rear] = element
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	element := q.buffer[q.front]
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front++
	}
	return element, nil
}

func (q *Queue) IsEmpty() bool {
	if q.front == -1 && q.rear == -1 {
		return true
	}
	return false
}

func (q *Queue) IsFull() bool {
	if q.front == 0 && q.rear-q.front == q.capacity-1 {
		return true
	}
	return false

}
func (*Queue) Peek() {

}
func (q *Queue) String() string {
	return fmt.Sprint(q.buffer[q.front:q.rear+1])
}
