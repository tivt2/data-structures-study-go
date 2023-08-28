package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []interface{}
}

func NewStack(options... int) (*Stack) {
	initialCapacity := 10
	if len(options) > 0 {
		initialCapacity = options[0]
	}

	return &Stack{
		data: make([]interface{}, 0, initialCapacity),
	}
}

func (s *Stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	length := len(s.data)
	out := s.data[length-1]
	s.data = s.data[:length-1]
	return out, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	return s.data[len(s.data)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Length() int {
	return len(s.data)
}

func (s *Stack) Print() {
	fmt.Println(s.data)
}

