package bitree

import (
	"github.com/vinthuy/mathlearn/tree"
)

//题目：输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求不能创建任何新的结点，只能调整树中结点指针的指向

//在二叉树中，每个结点都有两个指向子结点的指针。在双向链表中，每个结点也有两个指针，它们分别指向前一个结点和后一个结点。由于这两种结点的结构相似，
// 同时二叉搜索树也是一种排序的数据结构，因此在理论上有可能实现二叉搜索树和排序的双向链表的转换。在搜索二叉树中，左子结点的值总是小于父结点的值，
// 右子结点的值总是大于父结点的值。因此我们在转换成排序双向链表时，原先指向左子结点的指针调整为链表中指向前一个结点的指针，
// 原先指向右子结点的指针调整为链表中指向后一个结点指针。接下来我们考虑该如何转换。
//
//由于要求转换之后的链表是排好序的，我们可以中序遍历树中的每一个结点， 这是因为中序遍历算法的特点是按照从小到大的顺序遍历二叉树的每一个结点。
// 当遍历到根结点的时候，我们把树看成三部分：值为10 的结点、根结点值为6 的左子树、根结点值为1 4 的右子树。
// 根据排序链表的定义，值为10 的结点将和它的左子树的最大一个结点（即值为8 的结点）链接起来，同时它还将和右子树最小的结点（即值为12 的结点）链接起来，如图4.13 所示。

//  1       4->2->5->1->6-3->7
// 2   3
//4 5 6  7
func ToLinkList(bitree *tree.BiTree) *tree.BiTreeNode {
	//1.找到头部结点
	last := make([]*tree.BiTreeNode, 0, 1)
	last = append(last, nil)
	convertNode(bitree.Root, last)
	head := last[0]
	for head != nil && head.Left != nil {
		head = head.Left
	}
	return head
}

//中序遍历递归
func convertNode(curNode *tree.BiTreeNode, lastNode []*tree.BiTreeNode) {
	if curNode != nil {
		if curNode.Left != nil {
			convertNode(curNode.Left, lastNode)
		}
		if lastNode[0] != nil {
			lastNode[0].Right = curNode
			curNode.Left = lastNode[0]
		}
		lastNode[0] = curNode

		if curNode.Right != nil {
			convertNode(curNode.Right, lastNode)
		}
	}
}
