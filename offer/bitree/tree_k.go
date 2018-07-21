package bitree

import "github.com/vinthuy/mathlearn/tree"

//题目：给定一棵二叉搜索树，请找出其中的第k大的结点/请找出其中的第k大的结点
//解题思路
//　　如果按照中序遍历的顺序遍历一棵二叉搜索树，遍历序列的数值是递增排序的。只需要用中序遍历算法遍历一棵二叉搜索树，就很容易找出它的第k大结点。

func Find_K(root *tree.BiTreeNode,k[1]int) *tree.BiTreeNode {
	 var nd *tree.BiTreeNode
	 if root!=nil && root.Left!=nil{
	 	nd = Find_K(root.Left,k)
	 }

	 if n== nil{
	 	if k[0] == 1{
	 		return root
		}else {
			k[0] --
		}
	 }

	 if nd ==nil && root!=nil && root.Right!=nil{
	 	nd = Find_K(root.Right,k)
	 }

	 return nd
}