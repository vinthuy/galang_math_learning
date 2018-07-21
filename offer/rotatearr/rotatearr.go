package rotatearr

//题目： 把一个数组最开始的若干个元素搬到数组的末尾， 我们称之数组的旋转。
// 输入一个递增排序的数组的一个旋转， 输出旋转数组的最小元素。
// 例如数组{3，4, 5, 1, 2 ｝为{1,2,3, 4，5}的一个旋转，该数组的最小值为1
func RotateArr(array []int, low, hight int) int {
	length := len(array)
	if low < 0 || low > length-1 {
		return -1
	}

	if hight < 0 || hight > length-1 {
		return -1
	}
	var mid int
	for low < hight {
		if hight-low == 1 {
			return hight
		}
		mid = (hight +low) / 2
		if array[mid] > array[low] {
			low = mid
		} else if array[mid] < array[hight] {
			hight = mid
		}
	}
	return -1
}
