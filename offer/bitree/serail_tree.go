package bitree

import (
	"github.com/vinthuy/mathlearn/tree"
	"github.com/vinthuy/mathlearn/stack"
)

//题目：请实现两个函数，分别用来序列化和反序列化二叉树。

//解题思路
//　　通过分析解决前面的面试题6.我们知道可以从前序遍历和中序遍历构造出一棵二叉树。受此启发，我们可以先把一棵二叉树序列化成一个前序遍历序列和一个中序序列，然后再反序列化时通过这两个序列重构出原二叉树。
//　　这个思路有两个缺点。一个缺点是该方法要求二叉树中不能用有数值重复的结点。另外只有当两个序列中所有数据都读l出后才能开始反序列化。如果两个遍历序列的数据是从一个流里读出来的，那就可能需要等较长的时间。
//　　实际上如果二叉树的序列化是从根结点开始的话，那么相应的反序列化在根结点的数值读出来的时候就可以开始了。因此我们可以根据前序遍历的顺序来序列化二叉树，
//因为前序遍历是从根结点开始的。当在遍历二叉树碰到NULL指针时，这些NULL指针序列化成一个特殊的字符（比如‘$’）。另外，结点的数值之间要用一个特殊字符（比如’,’）隔开。

func Serial(tree1 *tree.BiTree)[]interface{}{
	node1 := tree1.Root
	stack1:=new(stack.Stack)
	stack1.Push(node1)
	result:=make([]interface{},0,10)
	for !stack1.Empty(){
		s1:=stack1.Pop()
		f :=false
		if s1 !=nil {
			node2:= s1.(*tree.BiTreeNode)
			if node2 != nil{
				result = append(result,node2.Data)
				stack1.Push(node2.Left)
				stack1.Push(node2.Right)
				f = true
			}
		}
		if !f{
			result = append(result,nil)
		}
	}
	return  result
}

func Deserail(result[]interface{},idx int) *tree.BiTreeNode  {
	if  len(result) == 0 || len(result) <=idx {
		return nil
	}
	root:= tree.NewBiTreeNode(result[idx])
	root.Left = Deserail(result,2*idx+1)
	root.Right = Deserail(result,idx*2+2)
	return  root
}