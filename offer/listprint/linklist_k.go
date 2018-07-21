package listprint

import (
	"github.com/vinthuy/mathlearn/list"
	"github.com/vinthuy/mathlearn/node"
)

//题目：输入一个链表，输出该链表中倒数第k 个结点．为了符合大多数人的习惯，本题从1 开始计数，即链表的尾结点是倒数第1
//个结点．例如一个链表有6 个结点，从头结点开始它们的值依次是1 、2、3、4、5 、6。这个个链表的倒数第3 个结点是值为4 的结点．

func Linklist_k(linklist *list.List, k int) interface{} {
	if linklist.Len() < k || k < 1 {
		return nil
	}
	var p1, p2 *node.Node
	var i int
	for i, p1 = 0, linklist.Head; i < k; i++ {
		p1 = p1.Next
	}

	for p2 = linklist.Head; p1 != nil; {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2.Data
}
