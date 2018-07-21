package listprint

import (
	"testing"
	"github.com/vinthuy/mathlearn/tree"
)

//题目：请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，
//第二层按照从右到左的顺序打印，即第一行按照从左到右的顺序打印，第二层按照从右到左顺序打印，第三行再按照从左到右的顺序打印，其他以此类推。

//解题思路
//　　按之字形顺序打印二叉树需要两个栈。我们在打印某一行结点时，把下一层的子结点保存到相应的栈里。如果当前打印的是奇数层，
//   则先保存左子结点再保存右子结点到一个栈里；如果当前打印的是偶数层，则先保存右子结点再保存左子结点到第二个栈里。
func TestPrintZHi(t *testing.T) {
	btr := tree.BuildBiTreeFromArray(1, 2, 3, 4, 5, 7, 8, 9, 10, 11, 12, 13, 14)
	PrintZHi(btr.Root)
}

func TestMap(t *testing.T) {
	m := make(map[string]int,4)
	t.Log(len(m))
	_,has:=m["123"]
	t.Log(has)
}
