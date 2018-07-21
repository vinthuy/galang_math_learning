package bitree

import "testing"

//题目：输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则返回true。否则返回false。假设输入的数组的任意两个数字都互不相同

//解题思路：
//在后序遍历得到的序列中， 最后一个数字是树的根结点的值。数组中前面的数字可以分为两部分：
// 第一部分是左子树结点的值，它们都比根结点的值小：
// 第二部分是右子树结点的值，它们都比根结点的值大。
func TestVerifySequenceOfBST(t *testing.T) {
	data := []int{4, 6, 12, 8, 16, 14, 10}
	t.Log(VerifySequenceOfBST(data, 0, len(data)-1))
}
