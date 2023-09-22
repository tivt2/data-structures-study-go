package stack

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	data []T
}

func NewStack[T any](options ...int) *Stack[T] {
	initialCapacity := 10
	if len(options) > 0 {
		initialCapacity = options[0]
	}

	return &Stack[T]{
		data: make([]T, 0, initialCapacity),
	}
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var ans T
		return ans, errors.New("stack is empty")
	}
	length := len(s.data)
	out := s.data[length-1]
	s.data = s.data[:length-1]
	return out, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var ans T
		return ans, errors.New("stack is empty")
	}

	return s.data[len(s.data)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Length() int {
	return len(s.data)
}

func (s *Stack[T]) Print() {
	fmt.Println(s.data)
}
