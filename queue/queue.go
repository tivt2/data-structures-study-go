package queue

import (
	"errors"
	"fmt"
)

type Node struct {
	value interface{}
	next *Node
}

func newNode(item interface{}) *Node {
	return &Node{
		value: item,
		next: nil,
	}
}

var ErrEmptyQueue = errors.New("queue is empty")

type Queue struct {
	head *Node
	tail *Node
	length int
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		length: 0,
	}
}

func (q *Queue) Enqueue(item interface{}) {
	node := newNode(item)
	q.length++
	
	if q.tail == nil {
		q.head, q.tail = node, node
		return
	}

	q.tail.next, q.tail = node, node
}

func (q *Queue) Deque() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrEmptyQueue
	}
	
	out := q.head
	q.head = out.next
	out.next = nil
	q.length--
	return out.value, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrEmptyQueue
	}

	return q.head.value, nil
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) Print() {
	curr := q.head
	for curr != nil {
		fmt.Print(curr.value, " -> ")
		curr = curr.next
	}
	fmt.Print("\n")
}