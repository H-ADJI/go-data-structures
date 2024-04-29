package queue

import (
	"testing"
)

func TestCicularQueueIsFull(t *testing.T) {
	testCases := []struct {
		q      CircularQueue
		isFull bool
	}{
		{CircularQueue{buffer: [10]int{}, front: -1, rear: -1}, false},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4}, front: 0, rear: 9}, true},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4}, front: 2, rear: 1}, true},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4}, front: 5, rear: 2}, false},
	}
	for _, tc := range testCases {
		if tc.q.IsFull() != tc.isFull {
			t.Fatalf("isFull should be %v but got %v", tc.isFull, tc.q.IsFull())

		}
	}
}

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

func TestCircularEnqueue(t *testing.T) {
	testCases := []struct {
		q           CircularQueue
		shouldError bool
	}{
		{CircularQueue{buffer: [10]int{}, front: -1, rear: -1}, false},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4}, front: 0, rear: 3}, false},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4}, front: -1, rear: -1}, false},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, front: 5, rear: 4}, true},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, front: 5, rear: 2}, false},
		{CircularQueue{buffer: [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, front: 0, rear: 9}, true},
	}
	for i, tc := range testCases {
		err := tc.q.Enqueue(1)
		if err == nil {
			if tc.shouldError {
				t.Fatalf("Test Case[%d] : enqueueing into %v shouldn't be possible", i, tc.q)
			}

		}
	}
}

// func TestCircularDequeue(t *testing.T) {
// }
