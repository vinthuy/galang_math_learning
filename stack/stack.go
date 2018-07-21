package stack

import (
	"github.com/vinthuy/mathlearn/node"
	"bytes"
)

/**     top
  1<-2<-3
* stack指针类型
 */
type Stack struct {
	top  *node.Node
	size int
}

func NewStack() *Stack {
	return new(Stack)
}

func (stack *Stack) Push(v interface{}) {
	if  v == nil{
		return
	}
	node := node.New2(v, stack.top)
	node.Next = stack.top
	stack.top = node
	stack.size++
}

func (stack *Stack) Pop() interface{} {
	if stack.top != nil {
		result := stack.top
		stack.top = stack.top.Next
		result.Next = nil
		stack.size--
		return result.Data
	}
	return nil
}

func (stack *Stack) TopNode() *node.Node {
	return stack.top
}

func (stack *Stack) Empty() bool {
	return stack.top == nil || stack.size == 0
}

func (this *Stack) Size() int {
	return this.size
}

func (stack *Stack) Top() interface{} {
	if stack.top != nil {
		return stack.top.Data
	}
	return nil
}

func (stack *Stack) String() string {
	var buffer bytes.Buffer
	for s1 := stack.top; s1 != nil; s1 = s1.Next {
		buffer.WriteString(s1.String())
		buffer.WriteString(",")
	}
	byte := buffer.Bytes()
	return string(byte[0 : len(byte)-1])
}
