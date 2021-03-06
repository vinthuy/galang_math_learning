package array

import "testing"

//例子说明：
//如输入一个长度为9 的数组｛ 1, 2, 3, 2, 2, 2, 5, 4, 2｝。由于数字2在数组中出现了5 次，超过数组长度的一半，因此输出2 。
//解题思路：
//解法一：基于Partition 函数的O（n）算法
//
//数组中有一个数字出现的次数超过了数组长度的一半。如果把这个数组排序，那么排序之后位于数组中间的数字一定就是那个出现次数超过数组长度一半的数字。
// 也就是说，这个数字就是统计学上的中位数，即长度为n 的数组中第n/2 大的数字。
//这种算法是受快速排序算法的启发。在随机快速排序算法中，我们先在数组中随机选择一个数字，然后调整数组中数字的顺序，
// 使得比选中的数字小数字都排在它的左边，比选中的数字大的数字都排在它的右边。如果这个选中的数字的下标刚好是n/2，
// 那么这个数字就是数组的中位数。如果它的下标大于n/2 ，那么中位数应该位于它的左边，我们可以接着在它的左边部分的数组中查找。如果它的下标小于n/2，那么中位数应该位于它的右边，我们可以接着在它的右边部分的数组中查找。这是一个典型的递归过程。
//
//解法二：根据数组组特点找出O(n)的算法
//
//数组中有一个数字出现的次数超过数组长度的一半，也就是说它出现的次数比其他所有数字出现次数的和还要多。因此我们可以考虑在遍历数组的时候保存两个值：
// 一个是数组中的一个数字， 一个是次数。当我们遍历到下～个数字的时候，如果下一个数字和我们之前保存的数字相同，则次数加l ：
// 如果下一个数字和我们之前保存的数字，不同，则次数减1 。如果次数为霉，我们需要保存下一个数字，并把次数设为1 。
// 由于我们要找的数字出现的次数比其他所有数字出现的次数之和还要多，那么要找的数字肯定是最后一次把次数设为1 时对应的数字。
//
//本题采用第二种实现方式

func TestHalfCountNum(t *testing.T) {
	t.Log(HalfCountNum([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}
