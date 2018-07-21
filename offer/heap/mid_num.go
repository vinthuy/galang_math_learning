package heap

import (
	"container/heap"
	)

//题目：如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有值排序之后位于中间的数值。如果数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。
//
//
//解题思路
//　　由于数据是从一个数据流中读出来的，数据的数目随着时间的变化而增加。如果用一个数据容器来保存从流中读出来的数据，当有新的数据流中读出来时，这些数据就插入到数据容器中。这个数据容器用什么数据结构定义更合适呢？
//　　数组是最简单的容器。如果数组没有排序，可以用Partition函数找出数组中的中位数。在没有排序的数组中插入一个数字和找出中位数的时间复杂度是O(1)和O(n)。
//　　我们还可以往数组里插入新数据时让数组保持排序，这是由于可能要移动O(n)个数，因此需要O(n)时间才能完成插入操作。在已经排好序的数组中找出中位数是一个简单的操作，只需要O(1)时间即可完成。
//　　排序的链表时另外一个选择。我们需要O(n)时间才能在链表中找到合适的位置插入新的数据。如果定义两个指针指向链表的中间结点（如果链表的结点数目是奇数，
// 那么这两个指针指向同一个结点），那么可以在O（1）时间得出中位数。此时时间效率与及基于排序的数组的时间效率一样。
//　　二叉搜索树可以把插入新数据的平均时间降低到O(logn)。但是，当二叉搜索树极度不平衡从而看起来像一个排序的链表时，插入新数据的时间仍然是O(n)。
// 为了得到中位数，可以在二叉树结点中添加一个表示子树结点数目的字段。有了这个字段，可以在平均O(logn)时间得到中位数，但差情况仍然是O(n).
//　　为了避免二叉搜索树的最差情况，还可以利用平衡的二叉搜索树，即AVL树。通常AVL树的平衡因子是左右子树的高度差。可以稍作修改，把AVL的平衡因
// 子改为左右子树结点数目只差。有了这个改动，可以用O(logn)时间往AVL树中添加一个新结点，同时用O(1)时间得到所有结点的中位数。
//　　AVL树的时间效率很高，但大部分编程语言的函数库中都没有实现这个数据结构。应聘者在短短几十分钟内实现AVL的插入操作是非常困难的。于是我们不得不再分析还有没有其它的方法。
//　　如果能够保证数据容器左边的数据都小于右边的数据，这样即使左、右两边内部的数据没有排序，也可以根据左边最大的数及右边最小的数得到中位数。
// 如何快速从一个容器中找出最大数？用最大堆实现这个数据容器，因为位于堆顶的就是最大的数据。同样，也可以快速从最小堆中找出最小数。　　
// 因此可以用如下思路来解决这个问题：用一个最大堆实现左边的数据容器，用最小堆实现右边的数据容器。往堆中插入一个数据的时间效率是O(logn).由于只需O(1)时间就可以得到位于堆顶的数据，因此得到中位数的时间效率是O(1).
//　　接下来考虑用最大堆和最小堆实现的一些细节。首先要保证数据平均分配到两个堆中，因此两个堆中数据的数目之差不能超过1（为了实现平均分配，可以在数据的总数目是偶数时把新数据插入到最小堆中，否则插入到最大堆中）。
//　　还要保证最大堆中里的所有数据都要小于最小堆中的数据。当数据的总数目是偶数时，按照前面分配的规则会把新的数据插入到最小堆中。如果此时新的数据比最大堆中的一些数据要小，怎么办呢？
//　　可以先把新的数据插入到最大堆中，接着把最大堆中的最大的数字拿出来插入到最小堆中。由于最终插入到最小堆的数字是原最大堆中最大的数字，这样就保证了最小堆中的所有数字都大于最大堆中的数字。
//当需要把一个数据插入到最大堆中，但这个数据小于最小堆里的一些数据时，这个情形和前面类似

type IntHeap []int

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *IntHeap) First() int {
	return (*h)[0]
}


type DynamicArray struct {
   max,min *IntHeap
}

func NewDyamicArray() *DynamicArray {
	d:=new(DynamicArray)
	d.max = new(IntHeap)
	d.min = new(IntHeap)
	heap.Init(d.min)
	heap.Init(d.max)
	return d
}

func (this *DynamicArray) insert(num int){
   if (this.min.Len() +this.max.Len())%2==0{
   	   if this.max.Len()>0 &&num<this.max.First(){
   	   	  heap.Push(this.max,num)
   	   	  num = heap.Pop(this.max).(int)
	   }
	   heap.Push(this.min,num)
   }else {
   		if this.min.Len()>0 && num>=this.min.First(){
			heap.Push(this.min,num)
			num = heap.Pop(this.min).(int)
		}
	   heap.Push(this.max,num)
	}
}

func (this *DynamicArray) GetMid() int {
	size:= this.max.Len()+this.min.Len()
	if size == 0 {
		panic("valid")
	}
	if size %2==1{
		return this.min.First()
	}else {
		return  (this.min.First() + this.max.First())/2
	}
}