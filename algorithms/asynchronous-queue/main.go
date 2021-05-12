package main

import (
	"errors"
	"sync"
)

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type QueueBackend struct {
	head *Node
	tail *Node

	// Keep track of the current Queue size
	size    uint32
	maxSize uint32
}

type ConcurrentQueue struct {
	lock *sync.Mutex

	// empty and full locks
	notEmpty *sync.Cond
	notFull  *sync.Cond

	// queue storage
	backend *QueueBackend
}

func (queue *QueueBackend) createNode(data interface{}) *Node {
	node := Node{}
	node.data = data
	node.next = nil
	node.prev = nil

	return &node
}

func (queue *QueueBackend) push(data interface{}) error {
	if queue.size >= queue.maxSize {
		err := errors.New("queue is full")
		return err
	}

	if queue.size == 0 {
		// creates a Genesis node
		node := queue.createNode(data)
		queue.head = node
		queue.tail = node

		queue.size++

		return nil
	}

	currentHeadNode := queue.head
	newHeadNode := queue.createNode(data)
	newHeadNode.next = currentHeadNode
	currentHeadNode.prev = newHeadNode

	queue.head = currentHeadNode
	queue.size++

	return nil
}

func (queue *QueueBackend) pop() (interface{}, error) {
	if queue.size == 0 {
		err := errors.New("The queue is empty")
		return nil, err
	}

	currentEnd := queue.tail
	newEnd := currentEnd.prev

	if newEnd != nil {
		newEnd.next = nil
	}

	queue.size--
	if queue.size == 0 {
		queue.head = nil
		queue.tail = nil
	}

	return currentEnd.data, nil
}

func (queue *QueueBackend) isEmpty() bool {
	return queue.size == 0
}

func (queue *QueueBackend) isFull() bool {
	return queue.size >= queue.maxSize
}

func (c *ConcurrentQueue) enqueue(data interface{}) error {
	c.lock.Lock()

	for c.backend.isFull() {
		// wait for empty
		c.notFull.Wait()
	}

	// insert into the Queue 
	err := c.backend.push(data)

	c.notEmpty.Signal()

	c.lock.Unlock()

	return err
}

func (c *ConcurrentQueue) dequeue() (interface{}, error) {
	c.lock.Lock()

	for c.backend.isEmpty() {
		c.notEmpty.Wait()
	}

	data, err := c.backend.pop()

	// signal notFull
	c.notFull.Signal()

	c.lock.Unlock()

	return data, err
}

func(c *ConcurrentQueue) getSize() uint32 {
	c.lock.Lock()
	size := c.backend.size
	c.lock.Unlock()

	return size
}

// NewConCurrentQueue creates a new Queue
func NewConcurrentQueue(maxSize uint32) *ConcurrentQueue {
	queue := ConcurrentQueue{}

	// Initialize the mutexes
	queue.lock =  &sync.Mutex{}
	queue.notFull = sync.NewCond(queue.lock)
	queue.notEmpty = sync.NewCond(queue.lock)

	// Initialize backend queue
	queue.backend = &QueueBackend{}
	queue.backend.size = 0
	queue.backend.head = nil
	queue.backend.tail = nil

	queue.backend.maxSize = maxSize
	return &queue
}
