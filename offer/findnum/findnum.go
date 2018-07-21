package findnum

//题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都
// 按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
// 从最左下角开始寻找
// 1,2,3
// 4,5,6
// 7,8,9
// 此时如果需要找到5 ，从 7 开始 ，判断是大于 row--，定位到行，列从左->右寻找方可
func FindItemInMatrix(matrix [][]int, number int) (bool, int, int) {
	rows := len(matrix) - 1
	cols := 0
	colsLen := len(matrix[0])

	for {
		if rows < 0 || cols < 0 || cols >= colsLen {
			break
		}
		if matrix[rows][cols] == number {
			return true, rows, cols
		}
		if matrix[rows][cols] > number {
			rows--
		} else {
			cols ++
		}
	}
	return false, -1, -1
}
