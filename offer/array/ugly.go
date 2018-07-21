package array

func IsUgly(num int) bool {
	for num%2 == 0 {
		num = num / 2
	}
	for num%3 == 0 {
		num = num / 3
	}
	for num%5 == 0 {
		num = num / 5
	}
	if num == 1 {
		return true
	}
	return false
}

//题目：我们把只包含因子2、3 和5 的数称作丑数（Ugly Number）。求从小到大的顺序的第1500个丑数。
//举例说明：
//例如6、8 都是丑数，但14 不是，它包含因子7。习惯上我们把1 当做第一个丑数。
//
//解题思路：
//第一种：逐个判断每个数字是不是丑数的解法，直观但不够高效。
//
//第二种：创建数组保存已经找到丑数，用空间换时间的解法。
//
//根据丑数的定义， 丑数应该是另一个丑数乘以2、3 或者5 的结果（1除外）。因此我们可以创建一个数组，里面的数字是排好序的丑数，每一个丑数都是前面的丑数乘以2、3或者5得到的。
//
//这种思路的关键在于怎样确保数组里面的丑数是排好序的。假设数组中已经有若干个丑数排好序后存放在数组中，并且把己有最大的丑数记做M，我们接下来分析如何生成下一个丑数。
// 该丑数肯定是前面某一个丑数乘以2、3 或者5 的结果， 所以我们首先考虑把已有的每个丑数乘以2。在乘以2 的时钝能得到若干个小于或等于M 的结果。
// 由于是按照顺序生成的，小于或者等于M 肯定己经在数组中了，我们不需再次考虑：还会得到若干个大于M 的结果，但我们只需要第一个大于M 的结果，
// 因为我们希望丑数是按从小到大的顺序生成的，其他更大的结果以后再说。我们把得到的第一个乘以2 后大于M 的结果记为M2，
// 同样，我们把已有的每一个丑数乘以3 和5，能得到第一个大于M 的结果M3 和M，那么下一个丑数应该是M2、M3 和M5这3个数的最小者。
//
//前面分析的时候，提到把已有的每个丑数分别都乘以2、3 和5。事实上这不是必须的，因为已有的丑数是按顺序存放在数组中的。
// 对乘以2而言， 肯定存在某一个丑数T2，排在它之前的每一个丑数乘以2 得到的结果都会小于已有最大的丑数，在它之后的每一个丑数乘以2 得到的结果都会太大。
// 我们只需记下这个丑数的位置， 同时每次生成新的丑数的时候，去更新这个T2。对乘以3 和5 而言， 也存在着同样的T3和T5。
//
//本题实现了两种方法
// (m5,m3,m2)[M]

func GetUglyNumber_Solution2(index int) int {
	if index <= 0 {
		return 0
	}

	pUglyNumbers := make([]int, index, index)
	pUglyNumbers[0] = 1
	nexUglyIndex := 1

	pUgly2, pUgly3, pUgly5 := pUglyNumbers[0], pUglyNumbers[0], pUglyNumbers[0]

	for nexUglyIndex < index {
		min := min3int(2*pUgly2, 3*pUgly3, 5*pUgly5)
		pUglyNumbers[nexUglyIndex] = min
		for pUgly2*2 <= min {
			pUgly2++
		}
		for pUgly3*3 <= min {
			pUgly3++
		}
		for pUgly5*5 <= min {
			pUgly5++
		}
		nexUglyIndex++
	}

	return pUglyNumbers[nexUglyIndex-1] //1不是丑数
}

func min3int(num1, num2, num3 int) int {
	var min int
	if num1 < num2 {
		min = num1
	} else {
		min = num2
	}

	if min > num3 {
		min = num3
	}
	return min
}
