package list

import (
	"testing"
	"github.com/vinthuy/mathlearn/list"
)

//题目：输入两个递增排序的链表，合并这两个链表并使新链表中的结点仍然是按照递增排序的
func TestLinkListMerge(t *testing.T) {
	list1 := list.NewList()
	list1.AddFromArray(1, 3, 5)

	list2 := list.NewList()
	list2.AddFromArray(2, 4,7,9)

	list3 := linkListMerge(list1, list2, func(v1 interface{}, v2 interface{}) int {
		a1 := v1.(int)
		a2 := v2.(int)
		if a1 > a2 {
			return 1
		} else if a1 < a2 {
			return -1
		}
		return 0
	})

	t.Log(list3)
}
