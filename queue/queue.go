package queue

import (
	"github.com/vinthuy/mathlearn/list"
)

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list.NewList()}
}

func (q *Queue) Enqueue(v interface{}) {
	q.list.Add(v)
}

func (q *Queue) Dequeue() interface{} {
	return q.list.RemoveHead()
}

func (q *Queue) Empty() {
	if q.list.Len() > 0 {
		q.list.Clear()
	}
}
