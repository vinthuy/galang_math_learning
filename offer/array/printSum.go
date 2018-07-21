package array

import "fmt"

//题目一：输入一个递增排序的数组和一个数字s，在数组中查找两个数，得它们的和正好是s。如果有多对数字的和等于s，输出任意一对即可。
//举例说明
//例如输入数组｛1 、2 、4、7 、11 、15 ｝和数字15.由于4+ 11 = 15 ，因此输出4 和11 。
//
//解题思路
//　　我们先在数组中选择两个数字，如果它们的和等于输入的s，我们就找到了要找的两个数字。如果和小于s 呢？我们希望两个数字的和再大一点。
//由于数组已经排好序了，我们可以考虑选择较小的数字后面的数字。因为排在后面的数字要大一些，那么两个数字的和也要大一些， 就有可能等于输入的数字s 了。
//同样， 当两个数字的和大于输入的数字的时候，我们可以选择较大数字前面的数字，因为排在数组前面的数字要小一些。
func SumCheck(arr []int, defaultSum int) {
	if len(arr) < 2 {
		panic("invalid arr for sum")
	}
	low, high := 0, len(arr)-1
	var sum int
	for low < high {
		sum = arr[low ] + arr[high]
		if sum > defaultSum {
			high --
		} else if sum < defaultSum {
			low ++
		} else {
			fmt.Println("low:", arr[low], "hight", arr[high])
			high --
			low++
		}
	}
}

//题目二：输入一个正数s，打印出所有和为s 的连续正数序列（至少两个数）。
//举例说明
//例如输入15，由于1+2+3+4+5=4＋5+6＝7+8=15，所以结果打出3 个连续序列1～5、4～6 和7～8。
//
//解题思路
//　　考虑用两个数small 和big 分别表示序列的最小值和最大值。首先把small 初始化为1, big 初始化为2。
//如果从small 到big 的序列的和大于s，我们可以从序列中去掉较小的值，也就是增大small 的值。如果从small 到big 的序列的和小于s，我们可以增大big，
//让这个序列包含更多的数字。因为这个序列至少要有两个数字，我们一直增加small 到(1+s)/2 为止。
//　　以求和为9 的所有连续序列为例，我们先把small 初始化为1, big 初始化为2。此时介于small 和big 之间的序列是｛1，2}，
//序列的和为3，小于9，所以我们下一步要让序列包含更多的数字。我们把big 增加1 变成3,此时序列为｛ I, 2，坷。
//由于序列的和是6，仍然小于9，我们接下来再增加big 变成4，介于small 和big 之间的序列也随之变成｛ l, 2, 3, 4｝。
//由于列的和10 大于9，我们要删去去序列中的一些数字， 于是我们增加small 变成2，此时得到的序列是｛2, 3, 4｝， 序列的和E好是9。
//我们找到了第一个和为9 的连续序列，把它打印出来。接下来我们再增加big，重复前面的过程，可以找到第二个和为9 的连续序列｛4，5}。
func PrintSubArr(defaultSum int) [][]int {
	result := make([][]int, 0, 5)
	low, high := 1, 2
	sum := low
	for low <= high && high < defaultSum {
		s1 := sum + high
		if s1 < defaultSum { //1,2,3,4,5
			sum += high
			high++
		} else if s1 > defaultSum {
			sum -= low
			low ++
		} else {
			result = append(result, []int{low, high})
			sum -= low
			low++
		}
	}
	return result
}
