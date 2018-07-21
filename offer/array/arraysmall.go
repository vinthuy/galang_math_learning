package array

import (
	"strings"
	"bytes"
)

//剑指Offer学习】【面试题33：把数组排成最小的数】
//例子说明：
//例如输入数组{3， 32, 321}，则扫描输出这3 个数字能排成的最小数字321323。
//
//解题思路：
//第一种：直观解法
//
//先求出这个数组中所有数字的全排列，然后把每个排列拼起来，最后求出拼起来的数字的最大值。
//
//第二种：排序解法
//
//找到一个排序规则，数组根据这个规则排序之后能排成一个最小的数字。要确定排序规则，就要比较两个数字，也就是给出两个数字m 和n，
// 我们需要确定一个规则判断m 和n 哪个应该排在前面，而不是仅仅比较这两个数字的值哪个更大。
//
//根据题目的要求，两个数字m 和n能拼接成数字m和m。如果mn < nm，那么我们应该打印出m，也就是m 应该排在n 的前面，我们定义此时m 小于n：
// 反之，如果nm < mn，我们定义n小于m。如果mn=nm,m 等于n。在下文中，符号“<”、“>”及“＝”表示常规意义的数值的大小关系，而文字“大于”、“小于”、“等于”表示我们新定义的大小关系。
//
//接下来考虑怎么去拼接数字，即给出数字m和n，怎么得到数字m和m 并比较它们的大小。亘接用数值去计算不难办到，
// 但需要考虑到一个潜在的问题就是m 和n 都在int 能表达的范围内，但把它们拼起来的数字m 和m 用int 表示就有可能溢出了，所以这还是一个隐形的大数问题。
//
//一个非常直观的解决大数问题的方法就是把数字转换成字符串。另外，由于把数字m 和n 拼接起来得到m 和m，它们的位数肯定是相同的，
// 因此比较它们的大小只需要按照字符串大小的比较规则就可以了。
//
//本题采用第二种方法实现
func CompareString(str1 string, str2 string) int {
	s1 := str1 + str2
	s2 := str2 + str1
	return strings.Compare(s1, s2)
}

func quiteSort(arr []string, start, end int) {
	if start < end {
		left := start
		pivot := arr[start]
		for start < end {
			for start < end && CompareString(arr[end], pivot) >= 0 {
				end --
			}
			arr[start] = arr[end]
			for start < end && CompareString(arr[start], pivot) <= 0 {
				start ++
			}
			arr[end] = arr[start]
		}
		arr[end] = pivot
		quiteSort(arr, left, start-1)
		quiteSort(arr, start+1, end)
	}
}

/**
     * 题目：输入一个正整数数组，把数组里所有数字拼接起来排成一个数，
     * 打印能拼接出的所有数字中最小的一个。
     * @param array 输入的数组
     * @return 输出结果
     */
func printMinNumber(array []string) string {
	if array == nil || len(array) < 1 {
		panic("Array must contain value")
	}
	quiteSort(array, 0, len(array)-1)
	var byteBuffer bytes.Buffer
	for _, str := range array {
		byteBuffer.WriteString(str)
	}
	return byteBuffer.String()
}
