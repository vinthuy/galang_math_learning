package listprint

import (
	"github.com/vinthuy/mathlearn/list"
	"fmt"
	stack2 "github.com/vinthuy/mathlearn/stack"
	"github.com/vinthuy/mathlearn/node"
)

//题目：输入个链表的头结点，从尾到头反过来打印出每个结点的值。
// 1->2->3->4 print 4,3,2,1
func PrintList(list *list.List, printType int) {
	if list == nil || list.Length == 0 {
		fmt.Println("<nil>")
	}

	if printType == 1 {
		recursionPrint(list)
		return
	} else if printType == 2 {
		whilePrint(list)
		return
	}
	stackPrint(list)
}

func recursionPrint(list *list.List) {
	node := list.Head
	recursionPrint0(node)
}

func recursionPrint0(node *node.Node) {
	if node == nil { //到尾巴打印
		return
	}
	recursionPrint0(node.Next)
	fmt.Print(node.String(), "\t")
}

// 1.反转打印
// 2，借用stack
func whilePrint(list1 *list.List) {
	list := list1.ReverseList2()
	for node := list.Head; node != nil; {
		fmt.Print(node.String(), "\t")
		node = node.Next
	}
}

func stackPrint(list *list.List) {
	stack := stack2.NewStack()
	for node := list.Head; node != nil; node = node.Next {
		stack.Push(node)
	}

	for s := stack.Pop(); s != nil; {
		fmt.Print(s.(*node.Node).String(), "\t")
		s = stack.Pop()
	}
}
