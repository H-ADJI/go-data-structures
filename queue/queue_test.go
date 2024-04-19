package queue

import (
	"testing"
)

const bufferSize = 5

func TestNewQueue(t *testing.T) {
	staticQueue := NewQueue(bufferSize)
	if len(staticQueue.buffer) != bufferSize {
		t.Fatalf("Buffer size should be %d but got %d", bufferSize, len(staticQueue.buffer))
	}
}

func TestEnqueue(t *testing.T) {
	testCases := []struct {
		sourceArray []int
		isFull      bool
		shouldError bool
	}{
		{[]int{1, 2, 3, 4, 5}, true, false},
		{[]int{1, 2, 3, 4, 5, 5}, true, true},
		{[]int{1, 2, 3, 4, 5, 5, 6}, true, true},
		{[]int{1, 2, 3}, false, false},
		{[]int{}, false, false},
	}
	for _, tc := range testCases {
		q := NewQueue(bufferSize)
		err := q.Populate(tc.sourceArray)
		if err == nil && tc.shouldError {
			t.Fatalf("Enqueueing should error %v ", tc.sourceArray)
		}
		if q.IsFull() != tc.isFull {
			t.Fatalf("isFull should be %v but got %v", tc.isFull, q.IsFull())

		}
	}
}

func TestDequeue(t *testing.T) {
	testCases := []struct {
		sourceArray []int
		isEmpty     bool
		shouldError bool
	}{
		{[]int{1, 2, 3, 4, 5}, false, false},
		{[]int{1}, true, false},
		{[]int{}, true, true},
	}
	for _, tc := range testCases {
		q := NewQueue(bufferSize)
		q.Populate(tc.sourceArray)
		_, err := q.Dequeue()
		if err == nil && tc.shouldError {
			t.Fatalf("Dequeueing should error %v ", tc.sourceArray)
		}
		if q.IsEmpty() != tc.isEmpty {
			t.Fatalf("isFull should be %v but got %v", tc.isEmpty, q.IsEmpty())

		}
	}
}
