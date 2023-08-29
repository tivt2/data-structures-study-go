package heap

import (
	"errors"
	"fmt"
)

var ErrEmptyHeap = errors.New("heap is empty")

type CompareFn func(item1, item2 interface{}) bool // compareFn(a, b) -> true sends item 'a' closer to Pop

type Heap struct {
	data []interface{}
	compareFn CompareFn
	capacity int
}

func NewHeap(compareFn CompareFn) *Heap {
	return &Heap{
		data: make([]interface{}, 0, 1),
		compareFn: compareFn,
		capacity: 1,
	}
}

func (h *Heap) Insert(item interface{}) {
	length := h.Length()
	if length == h.capacity {
		h.grow()
	}

	h.data = append(h.data, item)
	h.heapfyUp(length)
}

func (h *Heap) Pop() (interface{}, error) {
	if h.IsEmpty() {
		return nil, ErrEmptyHeap
	}
	out := h.data[0]

	lastIdx := h.Length() - 1
	h.swap(0, lastIdx)
	h.data = h.data[:lastIdx]

	h.heapfyDown(0)
	return out, nil
}

func (h *Heap) Peek() (interface{}, error) {
	if h.IsEmpty() {
		return nil, ErrEmptyHeap
	}
	return nil, nil
}

func (h *Heap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap) Length() int {
	return len(h.data)
}

func (h *Heap) Print() {
	fmt.Println(h.data)
}

func (h *Heap) heapfyDown(idx int) {
	leftChildIdx := leftChild(idx)
	if leftChildIdx >= h.Length() {
		return
	}

	leftChildValue := h.data[leftChildIdx]
	rightChildIdx := rightChild(idx)
	rightChildValue := h.data[rightChildIdx]
	curr := h.data[idx]
	if (h.compareFn(leftChildValue,rightChildValue)) {
		if (h.compareFn(leftChildValue, curr)) {
			h.swap(idx, leftChildIdx)
			h.heapfyDown(leftChildIdx)
		}
	} else if (h.compareFn(rightChildValue, curr)) {
			h.swap(idx, rightChildIdx)
			h.heapfyDown(rightChildIdx)
	}
}

func (h *Heap) heapfyUp(idx int) {
	if idx == 0 {
		return
	}

	parentIdx := parent(idx)
	parentValue := h.data[parentIdx]
	curr := h.data[idx]

	if (h.compareFn(curr, parentValue)) {
		h.swap(idx, parentIdx)
		h.heapfyUp(parentIdx)
	}
}

func (h *Heap) swap(idx1, idx2 int) {
	h.data[idx1], h.data[idx2] = h.data[idx2], h.data[idx1]
}

func parent(idx int) int {
	return (idx-1)/2
}

func leftChild(idx int) int {
	return 2*idx + 1
}

func rightChild(idx int) int {
	return 2*idx + 2
}

func (h *Heap) grow() {
	h.capacity = h.capacity * 2 + 1
	newData := make([]interface{}, len(h.data), h.capacity)
	copy(newData, h.data)
	h.data = newData
}
