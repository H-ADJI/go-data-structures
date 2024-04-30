package stack

import "errors"

const stackSize = 5

type Stack struct {
	buffer [stackSize]int
	index  int
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.index = -1
	return stack
}
func (s *Stack) Push(el int) error {
	if s.IsFull() {
		return errors.New("the stack is full")
	}
	s.index++
	s.buffer[s.index] = el
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("the stack is empty")
	}
	el := s.buffer[s.index]
	s.index--
	return el, nil
}

func (s *Stack) IsFull() bool {
	return s.index == stackSize-1
}

func (s *Stack) IsEmpty() bool {
	return s.index == -1
}
