package findnum

import (
	"math"
	"github.com/vinthuy/mathlearn/array"
)

/**
	* 题目：在一个长度为n的数组里的所有数字都在0到n-1的范围内。数组中某些数字是重复的，
	* 但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
	* 例如，如果输入长度为7的数组{2,3,1,0,2,5,3}，那么对应的输出是重复的数字2或者。
    {1,
*/
func GetRepNum(nums []int) int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if i > num && num == nums[num] {
			return num
		}
		for num != i {
			//swap numbers[i], numbers[num]
			array.Swap(nums, i, nums[i])
			num = nums[i]
		}

	}
	return math.MinInt8
}

/**
题目：给定一个数组A[0,1,…,n-1],请构建一个数组B[0,1,…,n-1],其中B中的元素B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1],不能使用除法。
解题思路
例如：A[]={1,2,3}求B[]
B[0]=A[1]×A[2]=2×3=6
B[1]=A[0]×A[2]=1×3=3
B[2]=A[0]×A[1]=1×2=2

1.B[0]初始化为1，从下标i=1开始，先求出C[i]的值并放入B[i],即B[i]=C[i]=C[i-1]×A[i-1]，所以B[1]=B[1-1]×A[1-1]=B[0]×A[0]=1×1=1,i++

2.B[2]=B[2-1]×A[2-1]=B[1]×A[1]=1×2=2,i++超出长度停止循环

3.C[i]计算完毕求D[i],设置一个临时变量temp初始化为1

4.从后往前变量数组，LengthA=3初始化i=LengthA-2=1,结束条件为i>=0

5.第一次循环，temp=temp×A[i+1]=1×A[2]=3,计算出A中最后一个元素的值放入temp,temp相当于D[i]的值

6.因为之前的B[i]=C[i],所以让B[i]×D[i]就是要保存的结果，即B[i]=B[1]=B[1]×temp=1×3=3,i–=0

7.计算B[i]=B[0],temp上一步中的值是A[2],在这次循环中temp=temp×A[0+1]=A[2]×A[1]=3×2=6

8.B[i]=B[0]=B0]×temp=1×6=6,i–<0循环结束

所以B数组为{6,3,2}
 */
func Multiply(matrix1 []int) []int {
	len1 := len(matrix1)
	var matrix2 = make([]int, len1, len1)
	//第一步每个result[i]都等于于data[0]*data[1]...data[i-1]
	//当i=n-1时，此时result[n-1]的结果已经计算出来了【A】
	matrix2[0] = 1
	for i := 1; i < len1; i++ {
		matrix2[i] = matrix2[i-1] * matrix1[i-1]
	}
	// tmp保存data[n-1]*data[n-2]...data[i+1]的结果
	tmp := 1
	// 第二步求data[n-1]*data[n-2]...data[i+1]
	// 【A】result[n-1]的结果已经计算出来，所以从data.length-2开始操作
	for i := len1 - 2; i >= 0; i-- {
		tmp *= matrix1[i+1]
		matrix2[i] *= tmp
	}
	return matrix2
}
