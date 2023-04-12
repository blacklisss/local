package data_structures

// Queue is a basic implementation of a queue data structure.
type Queue struct {
	items []interface{}
}

// NewQueue creates a new Queue instance.
func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

// Enqueue adds an item to the back of the queue.
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item at the front of the queue.
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Peek returns the item at the front of the queue without removing it.
func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.items[0]
}

// IsEmpty returns true if the queue is empty, false otherwise.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue.
func (q *Queue) Size() int {
	return len(q.items)
}
