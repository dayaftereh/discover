package container

import (
	"container/list"
	"sync"
)

// Queue is a simple concurrent queue
type Queue struct {
	l    *list.List
	lock sync.Mutex
}

// NewQueue creates a new queue
func NewQueue() *Queue {
	return &Queue{
		l: list.New(),
	}
}

// Push to the queue
func (queue *Queue) Push(v interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.l.PushBack(v)
}

// Len returns the item count in the queue
func (queue *Queue) Len() int {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	return queue.l.Len()
}

// Pop remove the head of the queue
func (queue *Queue) Pop() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	head := queue.l.Front()
	v := queue.l.Remove(head)
	return v
}
