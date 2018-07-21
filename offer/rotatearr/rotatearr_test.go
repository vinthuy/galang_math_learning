package rotatearr

import "testing"

//题目： 把一个数组最开始的若干个元素搬到数组的末尾， 我们称之数组的旋转。
// 输入一个递增排序的数组的一个旋转， 输出旋转数组的最小元素。
// 例如数组{3，4, 5, 1, 2 ｝为{1,2,3, 4，5}的一个旋转，该数组的最小值为1
func TestRotateArr(t *testing.T) {
	a := []int{3, 4, 1, 2}
	t.Log(RotateArr(a, 0, len(a)-1))
}