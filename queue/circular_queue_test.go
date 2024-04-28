package queue

import (
	"testing"
)

// func TestCicularQueueIsFull(t *testing.T) {

// testCases := []struct {
// 	buffer [QueueCapacity]int

// 	isFull bool
// }{
// 	{[QueueCapacity]int{1, 2, 3, 4, 5}, false},
// 	{[QueueCapacity]int{1, 2, 3, 4, 5, 5}, false},
// 	{[QueueCapacity]int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}, true},
// 	{[QueueCapacity]int{}, false},
// }
// for _, tc := range testCases {
// 	q := CircularQueue{}
// 	q.buffer = tc.buffer
// 	if q.IsFull() != tc.isFull {
// 		t.Fatalf("isFull should be %v but got %v", tc.isFull, q.IsFull())

// 	}
// }
// }

func TestCicularQueueIsEmpty(t *testing.T) {
	testCases := []struct {
		q       CircularQueue
		isEmpty bool
	}{
		{CircularQueue{buffer: [10]int{}, front: -1, rear: -1}, true},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4}, front: 0, rear: 3}, false},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4}, front: -1, rear: -1}, true},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4}, front: 5, rear: 2}, false},
	}
	for _, tc := range testCases {
		if tc.q.IsEmpty() != tc.isEmpty {
			t.Fatalf("isEmpty should be %v but got %v", tc.isEmpty, tc.q.IsEmpty())

		}
	}
}

func TestCicularQueueOccupiedSpace(t *testing.T) {
	testCases := []struct {
		q     CircularQueue
		space int
	}{
		{CircularQueue{buffer: [10]int{}, front: -1, rear: -1}, 0},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4}, front: 0, rear: 3}, 4},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4}, front: -1, rear: -1}, 0},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, front: 5, rear: 2}, 8},
	}
	for i, tc := range testCases {
		if tc.q.occupiedSpace() != tc.space {
			t.Fatalf("Test Case[%d] occupied space in the queue should be %v but got %v", i, tc.space, tc.q.occupiedSpace())

		}
	}
}

// func TestCircularEnqueue(t *testing.T) {
// }
// func TestCircularDequeue(t *testing.T) {
// }
