package queue

import "sync"

type ConcurrentQueue struct {
	queue []int32
	mutex sync.Mutex
}

func NewConcurrentQueue() *ConcurrentQueue {
	return &ConcurrentQueue{}
}

func (c *ConcurrentQueue) Enqueue(value int32) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.queue = append(c.queue, value)
}

func (c *ConcurrentQueue) Dequeue() int32 {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.queue) == 0 {
		panic("Queue is empty , and cannot be dequeued")
	}
	value := c.queue[0]
	c.queue = c.queue[1:]
	return value
}

func (c *ConcurrentQueue) IsEmpty() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return len(c.queue) == 0
}

func (c *ConcurrentQueue) Size() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return len(c.queue)
}
