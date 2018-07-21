package stack

import "testing"

func TestMinStack(t *testing.T) {
	minstack := NewMinStack()
	minstack.Push(2)
	minstack.Push(3)
	minstack.Push(7)
	minstack.Push(1)

	t.Log(minstack.Pop())
	t.Log(minstack.Pop())
	t.Log(minstack.Pop())
	t.Log(minstack.Pop())
}
