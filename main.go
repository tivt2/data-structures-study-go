package main

import (
	"fmt"
	// "root/heap"
	"root/lru"
	// "root/queue"
	// "root/ringBuffer"
	// "root/stack"
)

func main() {
	fmt.Println("running main")

	// stack := stack.NewStack()

	// stack.Push("item1")
	// stack.Push("item2")
	// out, _ := stack.Pop()
	// stack.Print()
	// fmt.Println(out)

	// queue := queue.NewQueue()

	// queue.Enqueue("value1")
	// queue.Enqueue("value2")
	// queue.Print()
	// out2, _ := queue.Deque()
	// fmt.Println(out2)

	// rb := ringBuffer.NewRingBuffer(2)

	// rb.Enqueue("item1")
	// rb.Print()
	// rb.Enqueue("item2")
	// rb.Print()
	// out3, _ := rb.Deque()
	// rb.Enqueue("item3")
	// rb.Print()
	// rb.Enqueue("item4")
	// rb.Print()
	// fmt.Println(out3)

	// compare := func(item1, item2 interface{}) bool {
	// 	return item1.(int) < item2.(int)
	// }

	// heap := heap.NewHeap(compare)

	// heap.Insert(1)
	// heap.Print()
	// heap.Insert(5)
	// heap.Print()
	// heap.Insert(2)
	// heap.Print()
	// heap.Insert(0)
	// heap.Print()
	// out, _ := heap.Pop()
	// fmt.Println(out)
	// heap.Print()

	// pubsub := pubSub.NewPubSub()

	// pubFn := func(text string) func(message interface{}) {
	// 	return func(message interface{}) {
	// 		fmt.Println(text, message.(string))
	// 	}
	// }

	// cancel := pubsub.Subscribe("something", pubFn("1"))
	// pubsub.Subscribe("something", pubFn("2"))
	// pubsub.Subscribe("another", pubFn("3"))
	// cancel()
	// pubsub.Publish("something", "publishing something")
	// pubsub.Publish("another", "publishing another")

	lru := lru.NewLRU(lru.Capacity(2))

	lru.Update("key1", "val1")
	lru.Print()
	lru.Update("key2", "val2")
	lru.Print()
	lru.Update("key3", "val3")
	lru.Print()
	lru.Get("key2")
	lru.Print()

}
