package bitree

import (
	"github.com/vinthuy/mathlearn/tree"
	"fmt"
)

type NodeInterface interface {
	equal(interface{}, interface{}) bool
}

//题目：输入两棵二叉树A 和B，判断B 是不是A 的子结构。

//解题思路：
//要查找树A 中是否存在和树B 结构一样的子树，我们可以分成两步：
// 第一步在树A 中找到和B 的根结点的值一样的结点R，
// 第二步再判断树A 中以R 为根结点的子树是不是包含和树B 一样的结构。

func HasSubTree(treeA *tree.BiTree, subTreeB *tree.BiTree, nodeInterface NodeInterface) bool {
	if treeA == nil || subTreeB == nil {
		return false
	}
	rootB := subTreeB.Root
	rootA := treeA.Root

	return findSubtree(rootA, rootB, nodeInterface)
}

func findSubtree(nodeA *tree.BiTreeNode, nodeB *tree.BiTreeNode, nodeInterface NodeInterface) bool {
	if nodeA == nil || nodeB == nil {
		return false
	}
	result := false

	if nodeInterface.equal(nodeA.Data, nodeB.Data) {
		fmt.Println("findSubtree", nodeA.Data, nodeB.Data)
		result = matchSubTree(nodeA, nodeB, nodeInterface)
	}

	if !result {
		result = findSubtree(nodeA.Left, nodeB, nodeInterface) || findSubtree(nodeA.Right, nodeB, nodeInterface)
	}
	return result
}

func matchSubTree(nodeA *tree.BiTreeNode, nodeB *tree.BiTreeNode, nodeInterface NodeInterface) bool {

	if nodeB == nil {
		return true
	}

	if nodeA == nil {
		return false
	}

	if nodeA == nodeB {
		return true
	}

	if nodeInterface.equal(nodeA.Data, nodeB.Data) {
		fmt.Println(nodeA.Data, nodeB.Data)
		return matchSubTree(nodeA.Left, nodeB.Left, nodeInterface) && matchSubTree(nodeA.Right, nodeB.Right, nodeInterface)
	}
	return false
}
