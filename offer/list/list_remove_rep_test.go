package list

import (
	"testing"
	"github.com/vinthuy/mathlearn/list"
)

//题目：在一个排序的链表中，如何删除重复的结点？
//解题思路
//　　解决这个问题的第一步是确定删除的参数。当然这个函数需要输入待删除链表的头结点。头结点可能与后面的结点重复，
//也就是说头结点也可能被删除，所以在链表头添加一个结点。
//　　接下来我们从头遍历整个链表。如果当前结点的值与下一个结点的值相同，那么它们就是重复的结点，都可以被删除。
//为了保证删除之后的链表仍然是相连的而没有中间断开，我们要把当前的前一个结点和后面值比当前结点的值要大的结点相连。
//我们要确保prev要始终与下一个没有重复的结点连接在一起。
// 1,1,1 2,2,3,3
func TestRemoveListRepeatNode(t *testing.T) {
	lst := list.NewList()
	lst.AddFromArray(0,1, 1, 1, 2, 2, 3, 3,4)
	t.Log(lst)
	RemoveListRepeatNode(lst)
	t.Log(lst)
}
