package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {

	dynamicQueue := NewQueue(-1)
	if len(dynamicQueue.buffer) > 0 {
		t.Fatal("Dynamic queue should have empty buffer")
	}
	// staticQueue := NewQueue(5)

}
