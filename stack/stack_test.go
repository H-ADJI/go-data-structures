package stack

import (
	"testing"
)

func TestStackIsFull(t *testing.T) {
	testCases := []struct {
		s      Stack
		isFull bool
	}{
		{Stack{buffer: [stackSize]int{}, index: -1}, false},
		{Stack{buffer: [stackSize]int{}, index: 2}, false},
		{Stack{buffer: [stackSize]int{}, index: stackSize - 1}, true},
		{Stack{buffer: [stackSize]int{1, 2, 3, 4, 5}, index: stackSize - 1}, true},
		{Stack{buffer: [stackSize]int{1, 2, 3, 4, 5}, index: -1}, false},
		{Stack{buffer: [stackSize]int{1, 2, 3, 4, 5}, index: 2}, false},
	}
	for i, tc := range testCases {
		if tc.s.IsFull() != tc.isFull {
			t.Fatalf("Test[%d] : isFull should be %v but got %v", i, tc.isFull, tc.s.IsFull())

		}
	}
}

func TestStackIsEmpty(t *testing.T) {
	testCases := []struct {
		s       Stack
		isEmpty bool
	}{
		{Stack{buffer: [stackSize]int{}, index: -1}, true},
		{Stack{buffer: [stackSize]int{}, index: 2}, false},
		{Stack{buffer: [stackSize]int{}, index: stackSize - 1}, false},
		{Stack{buffer: [stackSize]int{1, 2, 3, 4, 5}, index: -1}, true},
	}
	for i, tc := range testCases {
		if tc.s.IsEmpty() != tc.isEmpty {
			t.Fatalf("Test[%d] : IsEmpty should be %v but got %v", i, tc.isEmpty, tc.s.IsEmpty())

		}
	}

}
func TestStackPush(t *testing.T) {
	testCases := []struct {
		s           Stack
		shouldError bool
	}{
		{Stack{buffer: [stackSize]int{}, index: -1}, false},
		{Stack{buffer: [stackSize]int{}, index: 2}, false},
		{Stack{buffer: [stackSize]int{}, index: stackSize - 1}, true},
		{Stack{buffer: [stackSize]int{1, 2, 3, 4, 5}, index: -1}, false},
	}
	for i, tc := range testCases {
		err := tc.s.Push(1)
		if err == nil && tc.shouldError {
			t.Fatalf("Test[%d] : Shouldn't be able to push to stack %v", i, tc.s)

		}
	}

}
func TestStackPop(t *testing.T) {
	testCases := []struct {
		s           Stack
		shouldError bool
	}{
		{Stack{buffer: [stackSize]int{}, index: -1}, true},
		{Stack{buffer: [stackSize]int{}, index: 2}, false},
		{Stack{buffer: [stackSize]int{}, index: stackSize - 1}, false},
		{Stack{buffer: [stackSize]int{1, 2, 3, 4, 5}, index: -1}, true},
	}
	for i, tc := range testCases {
		_, err := tc.s.Pop()
		if err == nil && tc.shouldError {
			t.Fatalf("Test[%d] : Shouldn't be able to push to stack %v", i, tc.s)

		}
	}
}
