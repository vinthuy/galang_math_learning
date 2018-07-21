package bitree

import "github.com/vinthuy/mathlearn/tree"

//题目：请完成一个函数，输入一个二叉树，该函数输出它的镜像。
//解题思路：
// 我们先前序遍历这棵树的每个结点，如果遍历到的结点有子结点，就交换它的两个子结点。
// 当交换完所有非叶子结点的左右子结点之后，就得到了树的镜像。
func Mirror(btree *tree.BiTree) {
	root := btree.Root
	mirrorNode(root)
}

func mirrorNode(btreeNode *tree.BiTreeNode) {
	var tmpNode *tree.BiTreeNode
	if btreeNode.Left != nil || btreeNode.Right != nil {
		//交换左右结点
		tmpNode = btreeNode.Right
		btreeNode.Right = btreeNode.Left
		btreeNode.Left = tmpNode
	}
	if btreeNode.Left != nil {
		mirrorNode(btreeNode.Left)
	}
	if btreeNode.Right != nil {
		mirrorNode(btreeNode.Right)
	}
}
