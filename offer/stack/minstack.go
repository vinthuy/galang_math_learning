package stack

import (
	"github.com/vinthuy/mathlearn/stack"
	"math"
	"errors"
)

//题目： 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小素的min 函数。在该栈中，调用min、push 及pop的时间复杂度都是0(1)
//解题思路：
// 把每次的最小元素（之前的最小元素和新压入战的元素两者的较小值）都保存起来放到另外一个辅助栈里
// 如果每次都把最小元素压入辅助栈， 那么就能保证辅助栈的栈顶一直都是最小元素．
// 当最小元素从数据栈内被弹出之后，同时弹出辅助栈的栈顶元素，此时辅助栈的新栈顶元素就是下一个最小值。
type MinStack struct {
	min *stack.Stack
	s   *stack.Stack
}

func NewMinStack() *MinStack {
	return &MinStack{
		min: stack.NewStack(),
		s:   stack.NewStack(),
	}
}

func (this *MinStack) Push(v int) {
	this.s.Push(v)
	min := this.min.Top()
	var minNew int
	if min == nil {
		minNew = v
	} else {
		minNew = int(math.Min(float64(min.(int)), float64(v)))
	}
	this.min.Push(minNew)
}

func (this *MinStack) Pop() (error, int) {
	if this.s.Empty() {
		return errors.New("null queue"), math.MinInt8
	}
	this.s.Pop()
	return nil, this.min.Pop().(int)
}

func (this *MinStack) Min() (error, int) {
	if this.s.Empty() {
		return errors.New("null queue"), math.MinInt8
	}
	this.s.Pop()
	return nil, this.min.Top().(int)
}
