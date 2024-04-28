package queue

import "errors"

const QueueCapacity = 10

type CircularQueue struct {
	buffer [QueueCapacity]int
	front  int
	rear   int
}

func (q CircularQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("the queue is empty")
	}

	e := q.buffer[q.front]
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front + 1) % (QueueCapacity - 1)
	}
	return e, nil
}

func (q CircularQueue) Enqueue(e int) error {
	if q.IsFull() {
		return errors.New("the queue is full")
	}
	q.rear = (q.rear + 1) % (QueueCapacity - 1)
	q.buffer[q.rear] = e
	return nil
}

func (q CircularQueue) occupiedSpace() int {
	if q.IsEmpty() {
		return 0
	}
	if q.rear >= q.front {
		return q.rear - q.front + 1
	} else {
		return QueueCapacity + q.rear - q.front + 1
	}

}
func (q CircularQueue) IsFull() bool {
	return q.occupiedSpace() == QueueCapacity
}

func (q CircularQueue) IsEmpty() bool {
	if q.front == -1 && q.rear == -1 {
		return true
	}
	return false
}
func (q CircularQueue) Peek() (int, error) {
	return 0, nil
}
