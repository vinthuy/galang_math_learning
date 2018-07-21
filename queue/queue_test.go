package queue

import (
	"testing"
	"fmt"
)

func TestQueue(t *testing.T) {
	queue:=NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	queue.Enqueue(4)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
