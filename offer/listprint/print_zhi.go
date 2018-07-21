package listprint

import (
	"github.com/vinthuy/mathlearn/tree"
	"github.com/vinthuy/mathlearn/stack"
	"fmt"
)

//题目：请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，
//第二层按照从右到左的顺序打印，即第一行按照从左到右的顺序打印，第二层按照从右到左顺序打印，第三行再按照从左到右的顺序打印，其他以此类推。

//解题思路
//　　按之字形顺序打印二叉树需要两个栈。我们在打印某一行结点时，把下一层的子结点保存到相应的栈里。如果当前打印的是奇数层，
//   则先保存左子结点再保存右子结点到一个栈里；如果当前打印的是偶数层，则先保存右子结点再保存左子结点到第二个栈里。
func PrintZHi(root *tree.BiTreeNode) {
	if root == nil {
		return
	}
	s1 := stack.NewStack()
	s2 := stack.NewStack()
	s1.Push(root)
	var i = 0
	var node1 *tree.BiTreeNode
	for {
		if s1.Empty() && s2.Empty() {
			break
		}

		if i%2 == 0 {
			for {
				n := s1.Pop()
				if n == nil {
					break
				}
				node1 = n.(*tree.BiTreeNode)
				fmt.Println(node1)
				if node1 != nil {
					if node1.Right != nil {
						s2.Push(node1.Right)
					}
					if node1.Left != nil {
						s2.Push(node1.Left)
					}
				}
			}

		} else {
			for {
				n := s2.Pop()
				if n == nil {
					break
				}
				node1 = n.(*tree.BiTreeNode)
				fmt.Println(node1)
				if node1 != nil {
					if node1.Left != nil {
						s1.Push(node1.Left)
					}
					if node1.Right != nil {
						s1.Push(node1.Right)
					}
				}
			}
		}
		i++
	}

}
