package bitree

//题目：输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则返回true。否则返回false。假设输入的数组的任意两个数字都互不相同

//解题思路：
//在后序遍历得到的序列中， 最后一个数字是树的根结点的值。数组中前面的数字可以分为两部分：
// 第一部分是左子树结点的值，它们都比根结点的值小：
// 第二部分是右子树结点的值，它们都比根结点的值大。

func VerifySequenceOfBST(a []int, start, end int) bool {
	// 如果对应要处理的数据只有一个或者已经没有数据要处理（start>end）就返回true
	if start >= end {
		return true
	}

	m := a[end]
	// 找到左半部分[start,mid]
	mid := -1
	if start < end {
		for j := 0; a[j] < m; j++{
			mid = j
		}
		//只要出现 < 就 return false
		for i := mid + 1; i < end; i++ {
			for a[i] < m {
				return false
			}
		}
	}
	// 找到右半分[mid+1,end]
	return VerifySequenceOfBST(a, start, mid) && VerifySequenceOfBST(a, mid+1, end-1)
}
