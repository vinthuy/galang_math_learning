package bitree

import (
	"testing"
	"github.com/vinthuy/mathlearn/tree"
)

type IntNode struct {
}

func (this *IntNode) equal(v1 interface{}, v2 interface{}) bool {
	return v1.(int) == v2.(int)
}

//题目：输入两棵二叉树A 和B，判断B 是不是A 的子结构。

//解题思路：
//要查找树A 中是否存在和树B 结构一样的子树，我们可以分成两步：
// 第一步在树A 中找到和B 的根结点的值一样的结点R，
// 第二步再判断树A 中以R 为根结点的子树是不是包含和树B 一样的结构。

func TestHasSubTree(t *testing.T) {
	rootA := tree.NewBiTree(8)
	rootA.Root.Left = tree.NewBiTreeNode(8)
	rootA.Root.Right = tree.NewBiTreeNode(7)

	rootA.Root.Left.Left = tree.NewBiTreeNode(9)
	rootA.Root.Left.Right = tree.NewBiTreeNode(2)
	rootA.Root.Left.Right.Left = tree.NewBiTreeNode(4)
	rootA.Root.Left.Right.Right = tree.NewBiTreeNode(7)

	rootB := tree.NewBiTree(8)
	rootB.Root.Left = tree.NewBiTreeNode(9)
	rootB.Root.Right = tree.NewBiTreeNode(2)
	intNode := new(IntNode)
	t.Log(HasSubTree(rootA, rootB, intNode))
}
