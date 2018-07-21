package list

import (
	"github.com/vinthuy/mathlearn/list"
	"fmt"
)

//题目：0, 1, … , n-1 这n个数字排成一个圈圈，从数字0开始每次从圆圏里删除第m个数字。求出这个圈圈里剩下的最后一个数字。
//解题思路
//第一种：经典的解法， 用环形链表模拟圆圈。
//　　创建一个总共有n 个结点的环形链表，然后每次在这个链表中删除第m 个结点。
//0,1,2,3,4-->1,3,4->1,3
func ClRemove(n, m int) int {
	if n < 1 || m < 1 {
		return -1
	}
	lst := list.NewList()
	for i := 0; i < n; i++ {
		lst.Add(i)
	}

	var index int
	for lst.Len() > 1 {
		for i := 1; i < m; i++ {
			index = (index + 1) % lst.Len()
			fmt.Println("index", index)
		}
		lst.RemoveByIndex(index)
	}
	return lst.Head.Data.(int)
}

func ClRemove2(n, m int) int {
	if n < 1 || m < 1 {
		return -1
	}
	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}
