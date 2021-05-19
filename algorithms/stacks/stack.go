package stacks

import "sync"

type ItemType interface{}

type Stack struct {
	// items : holds items in the stack
	items  []ItemType
	rwlock sync.RWMutex
}

func (stack *Stack) Push(item ItemType) {
	if stack.items == nil {
		stack.items = []ItemType{}
	}

	stack.rwlock.Lock()
	stack.items = append(stack.items, item)
	stack.rwlock.Unlock()
}

func (stack *Stack) Pop() *ItemType {
	// check if stack is empty
	if len(stack.items) == 0 {
		return nil
	}

	stack.rwlock.Lock()
	// poping out items from the slice
	item := stack.items[len(stack.items)-1]
	// Removing the top element and adjusing the length
	stack.items = stack.items[0 : len(stack.items)-1 ]
	stack.rwlock.Unlock()
	return &item
}

func (stack *Stack) Size() int {
	stack.rwlock.RLock()
	defer stack.rwlock.RUnlock()

	return len(stack.items)
}

func (stack *Stack) All() []ItemType {
	stack.rwlock.RLock()
	defer stack.rwlock.RUnlock()

	return stack.items
}

func (stack *Stack) IsEmpty() bool {
	stack.rwlock.RLock()
	defer stack.rwlock.RUnlock()

	return len(stack.items) == 0
}