/*
Breadth first search implementation.
Author: Pavel Popov <pavelpopov@outlook.com>
*/
package main

import (
	"fmt"

	"github.com/schmooser/algo-007/graph"
)

func main() {
	//queue
	q1 := graph.QueueItem{1}
	q2 := graph.QueueItem{2}
	q3 := graph.QueueItem{3}
	var queue graph.Queue
	queue = make(graph.Queue, 0)
	queue = queue.Push(q1)
	queue = queue.Push(q2)
	queue = queue.Push(q3)
	fmt.Println(queue)
	queue, _ = queue.Pop()
	fmt.Println(queue)

	//stack
	s1 := graph.StackItem{1}
	s2 := graph.StackItem{2}
	s3 := graph.StackItem{3}
	var stack graph.Stack
	stack = make(graph.Stack, 0)
	stack = stack.Push(s1)
	stack = stack.Push(s2)
	stack = stack.Push(s3)
	fmt.Println(stack)
	stack, _ = stack.Pop()
	fmt.Println(stack)

}
