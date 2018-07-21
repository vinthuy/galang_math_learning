package findnum

import (
	"testing"
	"fmt"
)

//题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都
// 按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
// 从最左下角开始寻找
// 1,2,3
// 4,5,6
// 7,8,9
// 此时如果需要找到5 ，从 7 开始 ，判断是大于 row--，定位到行，列从左->右寻找方可
func TestFindItemInMatrix(t *testing.T) {
	var matrix = [][]int{{1, 2, 3,}, {4, 5, 6}, {7, 8, 9}}
	x, y, have := FindItemInMatrix(matrix, 5)
	fmt.Println("find item index", x, y, have)
}
