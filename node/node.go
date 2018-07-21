package node

import "strconv"

type Node struct {
	Data interface{}
	Next *Node
}

func (p *Node) String() string {
	return strconv.Itoa(p.Data.(int))
}

func IfNullVal(p *Node) interface{} {
	if p == nil {
		return 0
	}
	return p.Data
}

func New(data interface{}) *Node {
	return &Node{Data: data}
}

func New2(data interface{}, next *Node) *Node {
	return &Node{Data: data, Next: next}
}

func Equal(node1 *Node, node2 *Node, eqhandler func(node1 *Node, node2 *Node) bool) bool {
	if node1 == node2 {
		return true
	}
	if node1 != nil && node2 != nil {
		if eqhandler == nil {
			return node1.Data == node2.Data
		}
		return eqhandler(node1, node2)
	}
	return false
}
