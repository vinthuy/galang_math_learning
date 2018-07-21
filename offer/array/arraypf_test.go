package array

import (
	"testing"
)

// 输入数字n，按顺序打印出从1最大的n位十进制数。比如输入3，则打印出1、2、3 一直到最大的3位数即999。
func TestPrint1ToMaxOfNDigists(t *testing.T) {
	Print1ToMaxOfNDigists(3)
}

func TestOddEvenRecord(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	OddEvenRecord(a)
	t.Log(a)
}
