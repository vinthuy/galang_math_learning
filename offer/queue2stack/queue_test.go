package queue2stack

import (
	"testing"
	"fmt"
)

func TestQueue2(t *testing.T) {
	queue2 := NewQueue2()
	queue2.Enqueue(1)
	queue2.Enqueue(2)
	queue2.Enqueue(3)
	fmt.Println(queue2.Dequeue())
	fmt.Println(queue2.Dequeue())
	queue2.Enqueue(4)
	fmt.Println(queue2.Dequeue())
	fmt.Println(queue2.Dequeue())
}
