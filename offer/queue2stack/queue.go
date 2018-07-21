package queue2stack

import "github.com/vinthuy/mathlearn/stack"

type Queue2 struct {
	stack1 *stack.Stack
	stack2 *stack.Stack
}

func NewQueue2() *Queue2 {
	return &Queue2{stack.NewStack(), stack.NewStack()}
}

//入队列
func (q *Queue2) Enqueue(v interface{}) {
	q.stack1.Push(v)
}

//出队列
func (q *Queue2) Dequeue() interface{} {
	if q.stack2.Empty() {
		//把stack1 里面的stack全部移动到stack2
		for {
			v := q.stack1.Pop()
			if v == nil {
				break
			}
			q.stack2.Push(v)
		}
	}
	return q.stack2.Pop()
}
