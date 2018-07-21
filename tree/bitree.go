package tree

import (
	"strconv"
	queue2 "github.com/vinthuy/mathlearn/queue"
	"github.com/vinthuy/mathlearn/stack"
)

type BiTreeNode struct {
	Data  interface{}
	Left  *BiTreeNode
	Right *BiTreeNode
}

func (this *BiTreeNode) String() string {
	switch this.Data.(type) {
	case int:
		return strconv.Itoa(this.Data.(int))
	case string:
		return this.Data.(string)
	default:
		return "<not transfer>"
	}
}

func (this *BiTreeNode) IsLeaf() bool {
	return this.Left == nil && this.Right == nil
}

func NewBiTreeNode(data interface{}) *BiTreeNode {
	return &BiTreeNode{Data: data}
}

type BiTree struct {
	Root *BiTreeNode
}

func NewBiTree(rootData interface{}) *BiTree {
	return &BiTree{&BiTreeNode{Data: rootData}}
}

func BuildBiTreeFromArray(datas ... interface{}) *BiTree {
	if datas == nil || len(datas) == 0 {
		return nil
	}
	len := len(datas)
	var btree = new(BiTree)
	var treeArray = make([]*BiTreeNode, 0, len)
	for i, t := range datas {
		node := NewBiTreeNode(t)
		treeArray = append(treeArray, node)
		if i == 0 {
			btree.Root = treeArray[0]
			continue
		}
		//1.找到父类索引，如果是奇数就是左结点，否则是右结点
		pi := (i - 1) / 2
		pNode := treeArray[pi]
		if i%2 == 0 {
			pNode.Right = node
		} else {
			pNode.Left = node
		}
	}
	treeArray = nil //clear treeArray,wait for GC
	return btree
}

//广度优先
func (this *BiTree) BreadthSearchForWhile(handler func(v interface{})) {
	q := queue2.NewQueue()
	q.Enqueue(this.Root)
	for e := q.Dequeue(); e != nil; e = q.Dequeue() {
		left := e.(*BiTreeNode).Left
		if left != nil {
			q.Enqueue(left)
		}
		right := e.(*BiTreeNode).Left
		if right != nil {
			q.Enqueue(e.(*BiTreeNode).Right)
		}
		handler(e)
	}
}

type VisitOrderInterface interface {
	VisitNode(v interface{})
	VisitLeaf(v interface{})
}

/// 遍历接口
//  a
// b c
func StackPreOrder(b *BiTree, visitor VisitOrderInterface) {
	stack := stack.NewStack()
	node := b.Root
	for node != nil {
		for node != nil {
			stack.Push(node)
			visitor.VisitNode(node.Data)
			if node.IsLeaf() {
				visitor.VisitLeaf(node.Data)
			}
			node = node.Left
		}
		for {
			node, _ = stack.Pop().(*BiTreeNode)
			if node == nil {
				break
			} else if node != nil && node.Right != nil {
				node = node.Right //跳出循环，继续前面的逻辑
				break
			}
		}
	}
}

func StackMidOrder(b *BiTree, visitor VisitOrderInterface) {
	stack := stack.NewStack()
	node := b.Root
	for node != nil {
		for node != nil {
			stack.Push(node)
			node = node.Left
		}
		for {
			node, _ = stack.Pop().(*BiTreeNode)
			if node != nil {
				visitor.VisitNode(node.Data)
				if node.IsLeaf() {
					visitor.VisitLeaf(node.Data)
				}
			}
			if node == nil {
				break
			} else if node != nil && node.Right != nil {
				node = node.Right //跳出循环，继续前面的逻辑
				break
			}
		}
	}
}

func StackAfterOrder(b *BiTree, visitor VisitOrderInterface) {
	var visitMap map[*BiTreeNode]bool
	stack := stack.NewStack()
	node := b.Root
	for node != nil {
		for node != nil {
			stack.Push(node)
			node = node.Left
		}
		for {
			node, _ = stack.Top().(*BiTreeNode)
			if node == nil {
				break
			} else if node.Right == nil || !visitMap[node] {
				nodetmp := stack.Pop().(*BiTreeNode)
				visitor.VisitNode(nodetmp.Data)
				if nodetmp.IsLeaf() {
					visitor.VisitLeaf(node.Data)
				}
			} else {
				visitMap[node] = true
				node = node.Right
				break
			}
		}
	}
}
