package list

import "testing"

//题目：0, 1, … , n-1 这n个数字排成一个圈圈，从数字0开始每次从圆圏里删除第m个数字。求出这个圈圈里剩下的最后一个数字。
//解题思路
//第一种：经典的解法， 用环形链表模拟圆圈。
//　　创建一个总共有n 个结点的环形链表，然后每次在这个链表中删除第m 个结点。
func TestClRemove(t *testing.T) {
	t.Log(ClRemove(5,3))
	t.Log(ClRemove2(5,3))
}