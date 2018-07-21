package dp

import (
	"math"
	"fmt"
)

//
//题目：把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s 的所有可能的值出现的概率。
//解题思路
//解法一：基于通归求解，时间效率不够高。
// 先把n个骰子分为两堆：第一堆只有一个，另一个有n- 1 个。单独的那一个有可能出现从1 到6 的点数。我们需要计算从1 到6 的每一种点数和剩下的n-1 个骰子来计算点数和。
// 接下来把剩下的n-1个骰子还是分成两堆，第一堆只有一个， 第二堆有n-2 个。我们把上一轮那个单独骰子的点数和这一轮单独骰子的点数相加， 
// 再和剩下的n-2 个骰子来计算点数和。分析到这里，我们不难发现这是一种递归的思路，递归结束的条件就是最后只剩下一个骰子。
//　　我们可以定义一个长度为“6n-n+1 的数组， 和为s 的点数出现的次数保存到数组第s-n 个元素里。
//
//解法二：基于循环求解，时间性能好
//我们可以考虑用两个数组来存储骰子点数的每一个总数出现的次数。在一次循环中， 第一个数组中的第n 个数字表示骰子和为n 出现的次数
// 在下一循环中，我们加上一个新的骰子，此时和为n 的骰子出现的次数应该等于上一次循环中骰子点数和为n-1 、n-2 、n-3 、n-4, n-5 与n-6 的次数的总和，
// 所以我们把另一个数组的第n个数字设为前一个数组对应的第n-1 、n-2 、n-3 、n-4、n-5与n-6之和。

func PrintProbability2(number, max int) { //number 摔子个数，max ==6
	if number < 1 || max > 1 {
		return
	}
	var probabilities = make([][]int, 2, 2)
	for i := 0; i < 2; i++ {
		probabilities[i] = make([]int, number*max+1)
	}
	for i := 0; i < max*number+1; i++ {
		probabilities[0][i] = 0
		probabilities[1][i] = 0
	}

	// 标记当前要使用的是第0个数组还是第1个数组
	flag := 0

	// 抛出一个骰子时出现的各种情况
	for i := 1; i <= max; i++ {
		probabilities[flag][i] = 1
	}

	for k := 2; k <= number; k++ {
		for i := 0; i < k; i++ {
			probabilities[1-flag][i] = 0
		}
		// 抛出k个骰子，所有和的可能
		for i := k; i <= max*k; i++ {
			probabilities[1-flag][i] = 0

			// 每个骰子的出现的所有可能的点数
			// 统计出和为i的点数出现的次数 和N的值等于=上一个数组f(n) = f(n-1)+f(n-2)+..+f(n-{max})
			for j := 1; j <= i && j <= max; j++ {
				probabilities[1-flag][i] += probabilities[flag][i-j]
			}
		}
		flag = 1 - flag
	}

	total := math.Pow(float64(max), float64(number))
	for i := number; i <= max*number; i++ {
		ratio := float64(probabilities[flag][i]) / total
		fmt.Println(ratio)
	}
}

//题目：从扑克牌中随机抽5张牌，判断是不是一个顺子， 即这5张牌是不是连续的。2～10为数字本身， A为1。 J为11、Q为12、 为13。小王可以看成任意数字。
//解题思路
//　　我们可以把5张牌看成由5个数字组成的数组。大、小王是特殊的数字，我们不妨把它们都定义为0，这样就能和其他扑克牌区分开来了。
//　　接下来我们分析怎样判断5个数字是不是连续的，最直观的方法是把数组排序。值得注意的是，由于0可以当成任意数字，我们可以用0去补满数组中的空缺。如果排序之后的数组不是连续的，
// 即相邻的两个数字相隔若干个数字，但只要我们有足够的。可以补满这两个数字的空缺，这个数组实际上还是连续的。举个例子，数组排序之后为{0，1，3，4，5}
// 在1和3之间空缺了一个2，刚好我们有一个0，也就是我们可以把它当成2去填补这个空缺。
//　　于是我们需要做3 件事情： 首先把数组排序，再统计数组中0 的个数，最后统计排序之后的数组中相邻数字之间的空缺总数。如果空缺的总数小于或者等于0 的个数，
// 那么这个数组就是连续的：反之则不连续。
//　　最后，我们还需要注意一点： 如果数组中的非0 数字重复出现，则该数组不是连续的。换成扑克牌的描述方式就是如果一副牌里含有对子，则不可能是顺子。

func IsStraight(numbers []int) bool {
	if numbers == nil || len(numbers) == 0 {
		return false
	}

	fmt.Println(numbers)
	quiteSort(numbers, 0, len(numbers)-1)
	fmt.Println("sorted", numbers)

	numberOfZero, numberOfGap := 0, 0

	for _, num := range numbers {
		if num == 0 {
			numberOfZero++
		}
	}

	var gap int
	for i := 1; i < len(numbers)-1; i++ {
		gap = numbers[i] - numbers[i-1]
		if gap > 1 {
			numberOfGap += gap - 1
		}
	}

	return numberOfGap == numberOfZero
}

//快速排序
func quiteSort(numbers []int, lo, hi int) {
	if lo>hi{
		return
	}
	index := partition(numbers, lo, hi)
	fmt.Println(index)
	quiteSort(numbers, lo, index-1)
	quiteSort(numbers, index+1, hi)
}

func partition(array []int, lo, hi int) int {
	//固定的切分方式
	key := array[lo]
	for lo < hi {
		for array[hi] > key && hi > lo { //从后半部分向前扫描
			hi--
		}
		array[lo] = array[hi]
		for array[lo] <= key && hi > lo {
			lo++
		}
		array[hi] = array[lo]
	}
	array[lo] = key
	return lo
}
