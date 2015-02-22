/*
Some structures suitable for graphs.
Author: Pavel Popov <pavelpopov@outlook.com>
*/
package graph

type Edge struct {
	From, To int
}

type Vortex struct {
	X int
}

type Vortices []Vortex

type Edges []Edge

// Queue
type QueueItem struct {
	Label int
}

type Queue []QueueItem

// Push adds item to the queue.
func (queue Queue) Push(item QueueItem) Queue {
	return append(queue, item)
}

// Pop removes first item from queue and returns it.
func (queue Queue) Pop() (Queue, QueueItem) {
	return queue[1:], queue[0]
}

// Stack
type StackItem struct {
	Label int
}

type Stack []StackItem

// Push adds item to the queue.
func (stack Stack) Push(item StackItem) Stack {
	return append(stack, item)
}

// Pop removes first item from queue and returns it.
func (stack Stack) Pop() (Stack, StackItem) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}
