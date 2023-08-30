package lru

import (
	"errors"
	"fmt"
)

type node struct {
	value interface{}
	next  *node
	prev  *node
}

func newNode(value interface{}) *node {
	return &node{
		value: value,
		next:  nil,
		prev:  nil,
	}
}

type options struct {
	capacity int
}

type configFn func(*options)

func defaultOptions() options {
	return options{
		capacity: 10,
	}
}

func Capacity(capacity int) configFn {
	return func(o *options) {
		o.capacity = capacity
	}
}

type LRU struct {
	options
	length int

	head *node
	tail *node

	lookup         map[interface{}]*node
	reverse_lookup map[*node]interface{}
}

func NewLRU(config ...configFn) *LRU {
	options := defaultOptions()
	for _, fn := range config {
		fn(&options)
	}
	return &LRU{
		options:        options,
		length:         0,
		head:           nil,
		tail:           nil,
		lookup:         make(map[interface{}]*node, options.capacity),
		reverse_lookup: make(map[*node]interface{}, options.capacity),
	}
}

func (lru *LRU) Update(key, value interface{}) {
	if lru.IsEmpty() {
		lru.length++
		newNode := newNode(value)
		lru.head, lru.tail = newNode, newNode
		lru.lookup[key], lru.reverse_lookup[newNode] = newNode, key
		return
	}

	existingNode, ok := lru.lookup[key]
	if !ok {
		lru.length++
		newNode := newNode(value)
		lru.preprendNode(newNode)
		if lru.length > lru.capacity {
			lru.evict()
		}
		lru.lookup[key], lru.reverse_lookup[newNode] = newNode, key
		return
	}

	existingNode.value = value
	if existingNode == lru.tail {
		return
	}

	lru.detachAndPrependNode(existingNode)
}

func (lru *LRU) Get(key interface{}) (value interface{}, err error) {
	node, ok := lru.lookup[key]
	if !ok {
		return nil, errors.New("key not found")
	}

	lru.detachAndPrependNode(node)
	return node.value, nil
}

func (lru *LRU) IsEmpty() bool {
	return lru.length == 0
}

func (lru *LRU) Length() int {
	return lru.length
}

func (lru *LRU) Print() {
	curr := lru.tail
	for curr != nil {
		fmt.Print(curr.value, " -> ")
		curr = curr.next
	}
	fmt.Print("\n")
}

func (lru *LRU) preprendNode(node *node) {
	oldTail := lru.tail
	oldTail.prev, lru.tail, node.next = node, node, oldTail
}

func (lru *LRU) detachAndPrependNode(node *node) {
	if nodePrev := node.prev; nodePrev != nil {
		nodePrev.next = node.next
	} else {
		return
	}

	if nodeNext := node.next; nodeNext != nil {
		nodeNext.prev = node.prev
	} else {
		lru.head = node.prev
	}

	node.prev = nil
	lru.preprendNode(node)
}

func (lru *LRU) evict() {
	oldHead := lru.head
	lru.head, oldHead.prev = oldHead.prev, nil
	lru.head.next = nil

	key := lru.reverse_lookup[oldHead]
	delete(lru.lookup, key)
	delete(lru.reverse_lookup, oldHead)

	lru.length--
}
