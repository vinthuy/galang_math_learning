package list

import (
	"github.com/vinthuy/mathlearn/list"
	"github.com/vinthuy/mathlearn/node"
)

//题目：输入两个递增排序的链表，合并这两个链表并使新链表中的结点仍然是按照递增排序的
func linkListMerge(list1 *list.List, list2 *list.List, compareFunc func(v1 interface{}, v2 interface{}) int) *list.List {
	if list.IsEmpty(list1) {
		return list2
	}
	if list.IsEmpty(list2) {
		return list1
	}

	listNew := list.NewList()
	p1, p2 := list1.Head, list2.Head
	for {
		if p1 == nil || p2 == nil {
			break
		}
		if compareFunc(p1.Data, p2.Data) < 1 {
			listNew.Add(p1.Data)
			p1 = p1.Next
		} else {
			listNew.Add(p2.Data)
			p2 = p2.Next
		}
	}
	addRemainList(p1, listNew)
	addRemainList(p2, listNew)
	return listNew
}

func addRemainList(p1 *node.Node, listNew *list.List) {
	for p1 != nil {
		listNew.Add(p1.Data)
		p1 = p1.Next
	}
}
