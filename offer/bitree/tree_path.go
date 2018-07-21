package bitree

import (
	"github.com/vinthuy/mathlearn/tree"
	"fmt"
)

//题目：输入一棵二叉树和一个整数， 打印出二叉树中结点值的和为输入整数的所有路径。从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。
//
//解题思路：
//由于路径是从根结点出发到叶结点， 也就是说路径总是以根结点为起始点，因此我们首先需要遍历根结点。
// 在树的前序、中序、后序三种遍历方式中，只有前序遍历是首先访问根结点的。
//
//当用前序遍历的方式访问到某一结点时， 我们把该结点添加到路径上，并累加该结点的值。如果该结点为叶结点并且路径中结点值的和刚好等于输入的整数，
// 则当前的路径符合要求，我们把它打印出来。如果当前结点不是叶结点，则继续访问它的子结点。当前结点访问结束后，递归函数将自动回到它的父结点。
// 因此我们在函数退出之前要在路径上删除当前结点并减去当前结点的值，以确保返回父结点时路径刚好是从根结点到父结点的路径。
// 不难看出保存路径的数据结构实际上是一个枝， 因为路径要与递归调用状态一致， 而递归调用的本质就是一个压栈和出栈的过程。

func PathCheck(b *tree.BiTree, defaultSum int) {
	sum := 0
	path := make([]int, 0, 10)
	PreOrderRecursion(b.Root, defaultSum, sum, path)
}

func PreOrderRecursion(treeNode *tree.BiTreeNode, defaultSum int, sum int, path []int) {
	if treeNode == nil {
		return
	}
	sum = sum + treeNode.Data.(int)
	path = append(path, treeNode.Data.(int))
	if sum < defaultSum {
		PreOrderRecursion(treeNode.Left, defaultSum, sum, path)
		PreOrderRecursion(treeNode.Right, defaultSum, sum, path)
	} else if sum == defaultSum { // 如果当前和与期望的和相等
		// 当前结点是叶结点，则输出结果
		if treeNode.Left == nil && treeNode.Right == nil {
			fmt.Println(path)
		}
	}
	// 移除当前结点
	path = path[0 : len(path)-1]
}
