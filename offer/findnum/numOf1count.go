package findnum

import (
	"strconv"
	"fmt"
)

//题目：输入一个整数n求从1 到n这n个整数的十进制表示中1 出现的次数。

//题解思路：
//第一种：不夸虑时间效率的解法
//
//累加1 到n 中每个整数1 出现的次数。我们可以每次通过对10 求余数判断整数的个位数字是不是1 。如果这个数字大于10，除以10 之后再判断个位数字是不是1 。
//
//第二种：从数字规律着手明显提高时间效率的解法
//
//21345 作为例子来分析。我们把从1 到21345 的所有数字分为两段， 一段是从1 到1345，另一段是从1346 到21345。
//
//我们先看从1346 到21345 中1 出现的次数。1 的出现分为两种情况。首先分析1出现在最高位（本例中是万位）的情况。从1346 到21345 的数字中，
// 1出现在10000～19999 这10000 个数字的万位中， 一共出现了10000(10^4)个。
//
//值得注意的是， 并不是对所有5 位数而言在万位出现的次数都是10000
//个。对于万位是1 的数字比如输入12345, 1 只出现在10000～ 12345 的万位，出现的次数不是10^4 次，而是2346 次，
// 也就是除去最高数字之后剩下的数字再加上1 （即2345+1=2346 次）。
//
//接下来分析1出现在除最高位之外的其他四位数中的情况。例子中1346～21345 这20000 个数字中后4 位中1 出现的次数是2000 次。由于最
//高位是2，我们可以再把1346～21345 分成两段， 1346～11345 和1 1346～21345 。每一段剩下的4 位数字中， 选择其中一位是1 ，
// 其余三位可以在0～9 这10 个数字中任意选择，因此根据排列组合原则，总共出现的次数是2*10^3=2000
//
//至于从l 到1345 中1 出现的次数，我们就可以用递归求得了。这也是我们为什么要把1～21345 分成1～ 1 345 和1346～21345 两段的原因。
// 因为把21345 的最高位去掉就变成1345 ，便于我们采用递归的思路。
//
//本题采用第二种解法

func NumHasOneCount(num int) int {
	//21345
	//1.先转换为数组
	str := strconv.Itoa(num)
	var numarr = make([]int, 0, len(str))
	for _, s := range str {
		numarr = append(numarr, int(s)-'0')
	}

	return numberOf1(numarr, 0)
}

//21345
func numberOf1(arr []int, curIndex int) int {

	if arr == nil || curIndex >= len(arr) || curIndex < 0 {
		return 0
	}

	// 要处理的数字的位数
	first := arr[curIndex]
	length := len(arr) - curIndex
	if length == 1 && first == 0 {
		return 0
	}

	if length == 1 && first > 0 {
		return 1
	}

	numFirstDigist := 0
	// 如果最高位不是1，如21345，在[1236, 21345]中，最高位1出现的只在[10000, 19999]中，出现1的次数是10^4方个
	if first > 1 {
		numFirstDigist = powerBase10(length - 1)
		// 如果最高位是1，如12345，在[2346, 12345]中，最高位1出现的只在[10000, 12345]中，总计2345+1个
	} else if first == 1 {
		numFirstDigist = atoi(arr, curIndex+1) + 1
	}
	// numOtherDigits，是[1346, 21345]中，除了第一位之外（不看21345中的第一位2）的数位中的1的数目
	numOtherDigits := first * (length - 1) * powerBase10(length-2)
	fmt.Println("numOtherDigits", numOtherDigits)
	// numRecursive是1-1234中1的的数目
	numRecursive := numberOf1(arr, curIndex+1)
	fmt.Println("numRecursive", numOtherDigits)
	return numFirstDigist + numOtherDigits + numRecursive
}

/**
  * 将数字数组转换成数值，如{1, 2, 3, 4, 5}，i = 2，结果是345
  */
func atoi(numbers []int, i int) int {
	result := 0
	for j := i; j < len(numbers); j++ {
		result = result*10 + numbers[j]
	}
	return result;
}

func powerBase10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}
