package tree

import (
	"fmt"
	"github.com/vinthuy/mathlearn/stack"
)

//type ->string map func
//var typeStringMap = map[string]TypeHandlerFunc{
//	"string": func(a interface{}) string {
//		return a.(string)
//	},
//	"int": func(a interface{}) string {
//		return strconv.Itoa(a.(int))
//	},
//}

//func combination(m, n int) int {
//	if n > m-n {
//		n = m - n
//	}
//
//	c := 1
//	for i := 0; i < n; i++ {
//		c *= m - i
//		c /= i + 1
//	}
//
//	return c
//}
//sort interface to general type.
type DataInterface interface {
	Compare(b interface{}) int
	GetType() string
	String() string
}

//tree-node
type Node struct {
	data  DataInterface //结点的值
	left  *Node         //左结点
	right *Node         //右结点

	visited bool //右序遍历是否已经访问
}

func (node *Node) String() string {
	return node.data.String()
}

func NewNode(data DataInterface) *Node {
	return &Node{data: data}
}

//binaryTree define
type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) insert(node *Node) {

}

//type TypeHandler interface {
//	parse(a interface{}) string
//}
//
//type TypeHandlerFunc func(a interface{}) string
//
//func (f TypeHandlerFunc) parse(a interface{}) string {
//	return f(a)
//}
//根左右
func PreOrderRecursion(treeNode *Node) {
	if treeNode == nil {
		return
	}
	nodeStr := treeNode.String()
	fmt.Print(nodeStr, "\t")
	PreOrderRecursion(treeNode.left)
	PreOrderRecursion(treeNode.right)
}

//左根右
func MidOrderRecursion(treeNode *Node) {
	if treeNode == nil {
		return
	}
	MidOrderRecursion(treeNode.left)
	fmt.Print(treeNode.String(), "\t")
	MidOrderRecursion(treeNode.right)
}

//左右根
func AfterOrderRecursion(treeNode *Node) {
	if treeNode == nil {
		return
	}

	AfterOrderRecursion(treeNode.left)
	AfterOrderRecursion(treeNode.right)
	fmt.Print(treeNode.String(), "\t")
}

//  a
// b c
func stackPreOrder(b *BinaryTree) {
	stack := stack.NewStack()
	node := b.root
	for node != nil {
		for node != nil {
			stack.Push(node)
			fmt.Print(node.data.String(), "\t")
			node = node.left
		}
		for {
			node, _ = stack.Pop().(*Node)
			if node == nil {
				break
			} else if node != nil && node.right != nil {
				node = node.right //跳出循环，继续前面的逻辑
				break
			}
		}
	}
}

func stackMidOrder(b *BinaryTree) {
	stack := stack.NewStack()
	node := b.root
	for node != nil {
		for node != nil {
			stack.Push(node)
			node = node.left
		}
		for {
			node, _ = stack.Pop().(*Node)
			if node != nil {
				fmt.Print(node.data.String(), "\t")
			}
			if node == nil {
				break
			} else if node != nil && node.right != nil {
				node = node.right //跳出循环，继续前面的逻辑
				break
			}
		}
	}
}

func stackAfterOrder(b *BinaryTree) {
	stack := stack.NewStack()
	node := b.root
	for node != nil {
		for node != nil {
			stack.Push(node)
			node = node.left
		}
		for {
			node, _ = stack.Top().(*Node)
			if node == nil {
				break
			} else if node.right == nil || node.visited {
				fmt.Print(stack.Pop().(*Node).data.String(), "\t")
			} else {
				node.visited = true
				node = node.right
				break
			}
		}
	}
}
