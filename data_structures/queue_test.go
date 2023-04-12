package data_structures

import "testing"

func TestQueue(t *testing.T) {
	// Create a new queue
	queue := NewQueue()

	// Test IsEmpty for a new queue
	if !queue.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}

	// Test Enqueue and Size
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	if size := queue.Size(); size != 3 {
		t.Errorf("Expected queue size to be 3, but got %d", size)
	}

	// Test Peek
	if item := queue.Peek(); item != 1 {
		t.Errorf("Expected Peek to return 1, but got %v", item)
	}

	// Test Dequeue
	if item := queue.Dequeue(); item != 1 {
		t.Errorf("Expected Dequeue to return 1, but got %v", item)
	}
	if size := queue.Size(); size != 2 {
		t.Errorf("Expected queue size to be 2, but got %d", size)
	}

	// Test Dequeue until the queue is empty
	if item := queue.Dequeue(); item != 2 {
		t.Errorf("Expected Dequeue to return 2, but got %v", item)
	}
	if item := queue.Dequeue(); item != 3 {
		t.Errorf("Expected Dequeue to return 3, but got %v", item)
	}
	if !queue.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}

	// Test Dequeue for an empty queue
	if item := queue.Dequeue(); item != nil {
		t.Errorf("Expected Dequeue to return nil, but got %v", item)
	}

	// Test Peek for an empty queue
	if item := queue.Peek(); item != nil {
		t.Errorf("Expected Peek to return nil, but got %v", item)
	}
}
