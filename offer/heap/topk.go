package heap

import "container/heap"

//题目： 输入n个整数，找出其中最小的k个数。

//解题思路：
//解法一：O(n)时间算法，只有可以修改输入数组时可用。
//
//可以基于Partition函数来解决这个问题。如果基于数组的第k个数字来调整，使得比第k个数字小的所有数字都位于数组的左边，
// 比第k个数字大的所有数字都位于数组的右边。这样调整之后，位于数组中左边的k个数字就是最小的k 个数字（这k 个数字不一定是排序的〉。
//
//解法二： O（nlogk）的算法，精剧适合处理海量数据。
//先创建一个大小为k的数据容器来存储最小的k个数字，接下来我们每次从输入的n个整数中读入一个数．如果容器中已有的数字少于k个，则直接把这次读入的整数放入容器之中：如果容器中己有k 数字了，
// 也就是容器己满，此时我们不能再插入新的数字而只能替换已有的数字。找出这己有的k 个数中的最大值，然后1在这次待插入的整数和最大值进行比较。如果待插入的值比当前己有的最大值小，则用这个数替换当前已有的最大值：
// 如果待插入的值比当前已有的最大值还要大，那么这个数不可能是最小的k个整数之一，于是我们可以抛弃这个整数。
//
//因此当容器满了之后，我们要做3 件事情： 一是在k 个整数中找到最大数： 二是有可能在这个容器中删除最大数： 三是有可能要插入一个新的数字。
// 我们可以使用一个大顶堆在O(logk）时间内实现这三步操作。
//
//本题实现了两种方法

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *myHeap) First() int {
	return (*h)[0]
}

/**
借用堆实现topN
 */
func TopK(in []int, k int) []int {
	h := new(myHeap)
	heap.Init(h)
	for _, v := range in {
		if h.Len() <= k {
			heap.Push(h, v)
		} else if h.First() >= v {
			heap.Pop(h)
			heap.Push(h, v)
		}
	}
	return (*h)[0:k]
}
