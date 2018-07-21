package array

import (
	"fmt"
	"github.com/vinthuy/mathlearn/array"
)

// 输入数字n，按顺序打印出从1最大的n位十进制数。比如输入3，则打印出1、2、3 一直到最大的3位数即999。
func Print1ToMaxOfNDigists(n int) {
	if n < 0 {
		return
	}
	digistByte := make([]byte, n, n)
	for i := 0; i < 10; i++ {
		digistByte[0] = byte('0') + byte(i)
		print1ToMaxOfNDigistsRecursviely(digistByte, n, 0)
	}
}

func print1ToMaxOfNDigistsRecursviely(arr []byte, n int, index int) {
	if index == n-1 {
		printNumber(arr)
		return
	}

	for i := 0; i < 10; i++ {
		arr[index+1] = byte('0') + byte(i)
		print1ToMaxOfNDigistsRecursviely(arr, n, index+1)
	}

}

// 00001 -->打印1
func printNumber(str []byte) {
	var no0index = -1
	for i, s := range str {
		if byte(s) != '0' {
			no0index = i
			break
		}
	}
	if no0index == -1 {
		return
	}
	len := len(str)
	fmt.Println("str:", string(str[no0index:len]))
}

//题目:输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位予数组的后半部分。
//这个题目要求把奇数放在数组的前半部分， 偶数放在数组的后半部分，因此所有的奇数应该位于偶数的前面。
// 也就是说我们在扫描这个数组的时候， 如果发现有偶数出现在奇数的前面，我们可以交换它们的顺序，交换之后就符合要求了。
//因此我们可以维护两个指针，第一个指针初始化时指向数组的第一个数字，它只向后移动：
// 第二个指针初始化时指向数组的最后一个数字， 它只向前移动。在两个指针相遇之前，第一个指针总是位于第二个指针的前面。
// 如果第一个指针指向的数字是偶数，并且第二个指针指向的数字是奇数，我们就交换这两个数字。

//handler为钩子函数，为判断标准,做到灵活多用，符合标准放到后面，否咋在前面
func Record(arr []int, handler func(i int) bool) {
	i, j := 0, len(arr)-1 //i 低位，j高位
	for i < j {
		for !handler(i) {
			i++
		}
		for handler(j) {
			j--
		}

		if i < j {
			array.Swap(arr, i, j)
		}
	}
}

func OddEvenRecord(arr []int) {
	Record(arr, func(i int) bool {
		if arr[i]%2 == 0 {
			return true
		}
		return false
	})
}
