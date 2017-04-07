package cashier

import (
	"testing"
)

func TestNewWorkerQueue(t *testing.T) {
	expected := &workerQueue{head: nil, tail: nil}
	result := NewWorkerQueue()
	if result != expected {
		t.Fail()
	}
}
