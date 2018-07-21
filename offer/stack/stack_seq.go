package stack

import (
	"github.com/vinthuy/mathlearn/stack"
	"github.com/vinthuy/mathlearn/queue"
)

//题目：输入两个整数序列，第一个序列表示栈的压入顺序，请判断二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。
//解题思路：
//解决这个问题很直观的想法就是建立一个辅助栈，把输入的第一个序列中的数字依次压入该辅助栈，并按照第二个序列的顺序依次从该栈中弹出数字。
//判断一个序列是不是栈的弹出序列的规律：如果下一个弹出的数字刚好是栈顶数字，那么直接弹出。
// 如果下一个弹出的数字不在栈顶，我们把压栈序列中还没有入栈的数字压入辅助栈，直到把下一个需要弹出的数字压入栈顶为止。
// 如果所有的数字都压入栈了仍然没有找到下一个弹出的数字，那么该序列不可能是一个弹出序列。

/**
	* 输入两个整数序列，第一个序列表示栈的压入顺序，请判断二个序列是否为该栈的弹出顺序。
	* 假设压入栈的所有数字均不相等。例如序列1 、2、3 、4、5 是某栈压栈序列，
	* 序列4、5、3、2、1是该压栈序列对应的一个弹出序列，
	* 但4、3、5、1、2就不可能是该压棋序列的弹出序列。
	*/

func IsStackPopSeq(enterSeq [] int, popSeq []int) bool {
	q := queue.NewQueue()
	for _, i := range enterSeq {
		q.Enqueue(i)
	}

	s := stack.NewStack()
	for _, number := range popSeq {
		if s.Top() != number {
			for {
				el := q.Dequeue()
				if el == nil {
					break
				}
				s.Push(el)
				if el == number {
					break
				}
			}
		}
		if s.Top() != number {
			return false
		}
		s.Pop()
	}
	return true
}
