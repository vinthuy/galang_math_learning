package list

import (
	"bytes"
	"github.com/vinthuy/mathlearn/node"
	stack2 "github.com/vinthuy/mathlearn/stack"
)

//整数相加
func ListAdd(list1 *List, list2 *List) *List {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	result := &List{Length: 0}
	list1.reverseList()
	list2.reverseList()
	var p1, p2 = list1.Head, list2.Head
	var nexBit = 0
	for p1 != nil || p2 != nil {
		sum := node.IfNullVal(p1).(int) + node.IfNullVal(p2).(int) + nexBit
		if sum > 9 {
			nexBit = sum / 10
			sum = sum % 10
		}
		result.Add(sum)
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = nil
		}
		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = nil
		}
	}

	if result != nil {
		result.reverseList()
	}
	return result
}

func (list *List) reverseList() {
	if list == nil || list.Head == nil {
		return
	}
	tmpNode := list.Head
	var nodePre, next *node.Node
	for tmpNode != nil {
		next = tmpNode.Next
		tmpNode.Next = nodePre
		nodePre = tmpNode
		tmpNode = next
	}
	list.last = list.Head
	list.Head = nodePre
}

func (list *List) ReverseList2() *List {
	stack := stack2.NewStack()
	for n := list.Head; n != nil; n = n.Next {
		stack.Push(n.Data)
	}
	listNew := NewList()
	var n1 interface{}
	for {
		n1 = stack.Pop()
		if n1 == nil {
			break
		}
		listNew.Add(n1)
	}
	return listNew
}

type List struct {
	Head   *node.Node
	last   *node.Node
	Length int
}

func NewList() *List {
	return &List{Length: 0}
}

func (list *List) Len() int {
	return list.Length
}

func IsEmpty(list *List) bool {
	return list == nil || list.Len() == 0
}

func (list *List) Clear() {
	list.Head = nil
	list.last = nil
	list.Length = 0
}

func (list *List) AddFromArray(array ...interface{}) bool {
	if len(array) == 0 {
		return false
	}
	for _, arr := range array {
		list.Add(arr)
	}
	return true
}

func (list *List) Add(data interface{}) {
	if data == nil {
		return
	}
	newNode := node.New(data)
	if list.Head == nil {
		list.Head = newNode
		list.last = newNode
	} else {
		lastNode := list.last
		lastNode.Next = newNode
		list.last = newNode
	}
	list.Length++
}

func (list *List) addNode(newNode *node.Node) {
	if list.Head == nil {
		list.Head = newNode
		list.last = newNode
	} else {
		lastNode := list.last
		lastNode.Next = newNode
		list.last = newNode
	}
	list.Length++
}

func (list *List) RemoveHead() interface{} {
	head := list.Head
	if list.Remove(head) {
		return head.Data
	}
	return nil
}

func (list *List) Remove(src *node.Node) bool {
	if list.last == nil {
		return false
	}
	var nodePre, n *node.Node
	nodePre = nil
	n = list.Head
	for n != nil {
		if n == src {
			if nodePre != nil {
				nodePre.Next = n.Next
				if n.Next == nil {
					list.last = nodePre
				}
				n.Next = nil
			} else if nodePre == nil {
				list.Head = n.Next
				if n.Next == nil {
					list.last = nil
				}
				//第一个
				n.Next = nil
			}
			list.Length--
			return true
		}
		nodePre = n
		n = n.Next
	}
	return false
}

func (list *List) RemoveByIndex(idx int) *node.Node {
	i := 0
	var nodePre, n *node.Node
	nodePre = nil
	n = list.Head
	for n != nil {
		if i == idx {
			//remove
			if nodePre != nil {
				nodePre.Next = n.Next
				if n.Next == nil {
					list.last = nodePre
				}
				n.Next = nil
			} else if nodePre == nil {
				list.Head = n.Next
				if n.Next == nil {
					list.last = nil
				}
				//第一个
				n.Next = nil
			}
			list.Length--
			return n
		}
		nodePre = n
		n = n.Next
		i++
	}
	return nil
}

func (list *List) String() string {
	var buffer bytes.Buffer
	for node := list.Head; node != nil; node = node.Next {
		buffer.WriteString(node.String())
		buffer.WriteString(",")
	}
	bytes := buffer.Bytes()
	return string(bytes[0 : len(bytes)-1])
}

// test code view
//func main() {
//	list1 := newList()
//	list2 := newList()
//	list1.add(9)
//	list1.add(8)
//	list1.add(7)
//	list1.reverseList()
//
//	fmt.Println(list1.String())
//
//	list2.add(1)
//	list2.add(2)
//
//	list2.reverseList()
//	fmt.Println(list2.String())
//
//	fmt.Println(listAdd(list1, list2))
//}
