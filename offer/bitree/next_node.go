package bitree

import (
	"github.com/vinthuy/mathlearn/tree"
)

//题目：给定一棵二叉树和其中的一个结点，如何找出中序遍历顺序的下一个结点？树中的结点除了有两个分别指向左右子结点的指针以外，还有一个指向父节点的指针。
//解题思路
//　　1.如果一个结点有右子树，那么它的下一个结点就是它的右子树中的左子结点。也就是说右子结点出发一直沿着指向左子结点的指针，我们就能找到它的下一个结点。
//　　2.接着我们分析一个结点没有右子树的情形。如果结点是它父节点的左子结点，那么它的下一个结点就是它的父结点。
//　　3. 如果一个结点既没有右子树，并且它还是它父结点的右子结点，这种情形就比较复杂。我们可以沿着指向父节点的指针一直向上遍历，
//直到找到一个是它父结点的左子结点的结点。如果这样的结点存在，那么这个结点的父结点就是我们要找的下一个结点。
func NextNode(n *tree.BiPiTreeNode) *tree.BiPiTreeNode {
	var targetNode *tree.BiPiTreeNode
	//1
	if n.Right != nil {
		targetNode = n.Right
		for targetNode != nil {
			targetNode = targetNode.Left
		}
		return targetNode
	} else {
		parent := n.Parent
		//2
		if n == parent.Left {
			return parent
		} else {
			//3
			for parent != nil {
				parent = parent.Parent
				n = parent
				if n == parent.Left {
					return parent
				}
			}
		}
	}
	return nil
}
