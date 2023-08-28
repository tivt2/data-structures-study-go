package main

import (
	"fmt"
	"root/queue"
	"root/stack"
)

func main() {
	fmt.Println("running main")

	stack := stack.NewStack()

	stack.Push("item1")
	stack.Push("item2")
	out, _ := stack.Pop()
	stack.Print()
	fmt.Println(out)

	queue := queue.NewQueue()

	queue.Enqueue("value1")
	queue.Enqueue("value2")
	queue.Print()
	out2, _ := queue.Deque()
	fmt.Println(out2)
}