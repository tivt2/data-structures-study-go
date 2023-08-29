package ring_buffer

import (
	"errors"
	"fmt"
)

var ErrEmptyBuffer = errors.New("ring buffer is empty")

type RingBuffer struct {
	capacity int
	data []interface{}
	head int
	tail int
}

func NewRingBuffer(options... int) *RingBuffer {
	initialCapacity := 8
	if len(options) > 0 {
		initialCapacity = options[0]
	}

	return &RingBuffer{
		capacity: initialCapacity,
		data: make([]interface{}, initialCapacity),
		head: 0,
		tail: 0,
	}
}

func (rb *RingBuffer) Enqueue(item interface{}) {
	if rb.capacity == rb.Length() {
		rb.grow()
	}

	rb.data[rb.tail % rb.capacity] = item
	rb.tail++
}

func (rb *RingBuffer) Deque() (interface{}, error) {
	if rb.IsEmpty() {
		return nil, ErrEmptyBuffer
	}

	idx := rb.head % rb.capacity
	out := rb.data[idx]
	rb.data[idx] = nil
	rb.head++
	return out, nil
}

func (rb *RingBuffer) Peek() (interface{}, error) {
	if rb.IsEmpty() {
		return nil, ErrEmptyBuffer
	}

	return rb.data[rb.head], nil
}

func (rb *RingBuffer) Length() int {
	return rb.tail - rb.head
}

func (rb *RingBuffer) IsEmpty() bool {
	return rb.Length() == 0
}

func (rb *RingBuffer) Print() {
	fmt.Println(rb.data)
	fmt.Println("capacity:", rb.capacity)
}

func (rb *RingBuffer) grow() {
	oldCapacity := rb.capacity
	rb.capacity *= 2
	newData := make([]interface{}, rb.capacity)
	for i := 0; i < rb.Length(); i++ {
		newData[i] = rb.data[(i+rb.head) % oldCapacity]
	}
	rb.tail = oldCapacity
	rb.head = 0
	rb.data = newData
}