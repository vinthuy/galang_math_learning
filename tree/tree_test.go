package tree

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"errors"
)

// 单元测试
// 测试全局函数，以TestFunction命名
// 测试类成员函数，以TestClass_Function命名
//func TestCombination(t *testing.T) {
//	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
//	for _, unit := range []struct {
//		m        int
//		n        int
//		expected int
//	}{
//		{1, 0, 1},
//		{4, 1, 4},
//		{4, 2, 6},
//		{4, 3, 4},
//		{4, 4, 1},
//		{10, 1, 10},
//		{10, 3, 120},
//		{10, 7, 120},
//	} {
//		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
//		if actually := combination(unit.m, unit.n); actually != unit.expected {
//			t.Errorf("combination: [%v], actually: [%v]", unit, actually)
//		}
//	}
//}
//
////性能测试
//func BenchmarkCombination(b *testing.B) {
//	// b.N会根据函数的运行时间取一个合适的值
//	for i := 0; i < b.N; i++ {
//		combination(i+1, rand.Intn(i+1))
//	}
//}
//
//func BenchmarkCombinationParallel(b *testing.B)  {
//	// 测试一个对象或者函数在多线程的场景下面是否安全
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			m := rand.Intn(100) + 1
//			n := rand.Intn(m)
//			combination(m, n)
//		}
//	})
//}
type IntFace int

func (i IntFace) Compare(b interface{}) int {
	b1 := b.(int)
	a1 := int(i)
	fmt.Println(i)
	if a1 == b1 {
		return 0
	} else if a1 > b1 {
		return 1
	}
	return -1
}

func (i IntFace) GetType() string {
	return reflect.TypeOf(i).String()
}

func (i IntFace) String() string {
	return strconv.Itoa(int(i))
}

func FromArray(a []int) *BinaryTree {
	var btree = &BinaryTree{}
	var treeArray = make([]*Node, 0, len(a))
	for i, t := range a {
		intFace := IntFace(t)
		node := NewNode(intFace)
		treeArray = append(treeArray, node)
		if i == 0 {
			btree.root = treeArray[0]
			continue
		}
		//1.找到父类索引，如果是奇数就是左结点，否则是右结点
		pi := (i - 1) / 2
		pNode := treeArray[pi]
		if i%2 == 0 {
			pNode.right = node
		} else {
			pNode.left = node
		}
	}
	return btree
}

var a = [7]int{1, 2, 3, 4, 5, 6, 7}
var btree = FromArray(a[0:])

func TestA(t *testing.T) {
	fmt.Println("前序遍历:")
	PreOrderRecursion(btree.root)
	fmt.Println()
	fmt.Println("中序遍历:")
	MidOrderRecursion(btree.root)
	fmt.Println()
	fmt.Println("后序遍历:")
	AfterOrderRecursion(btree.root)
}

func Test_stackPreOrder(t *testing.T) {
	stackPreOrder(btree)
}

func Test_stackMidOrder(t *testing.T) {
	stackMidOrder(btree)
}

func Test_stackAfterOrder(t *testing.T) {
	stackAfterOrder(btree)
}

func Try(fun func(), handler func(interface{})) {
	defer func() { //finally
		if err := recover(); err != nil { //get err then handler err with function.
			handler(err)
		}
	}()
	fun()
}

func TestTry(t *testing.T) {
	Try(func() {
		panic(errors.New("foo")) //抛异常
	}, func(e interface{}) { //
		t.Error(e)
	})
}
