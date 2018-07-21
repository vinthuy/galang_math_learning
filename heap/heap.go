package heap

//最小堆，最大堆 表示数组类
type DataFace interface {
	Less(i, j int) bool //<
	Swap(i, j int)
	Len() int
	//add
	Push(v interface{})
	Pop() interface{}
}

func Push(heap DataFace, v interface{}) {
	heap.Push(v)
	up(heap, heap.Len()-1)
}

func Pop(heap DataFace) interface{} {
	if heap.Len() > 1 {
		heap.Swap(0, heap.Len()-1)
		down(heap, 0, heap.Len())
	}
	return heap.Pop()
}

func Remove(heap DataFace, i int) interface{} {
	if heap.Len() > 1 {
		n := heap.Len() - 1
		heap.Swap(i, n)
		if !down(heap, i, n) {
			up(heap, i)
		}
	}
	return heap.Pop()
}

func Fix(h DataFace, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

//以最大堆方式
func up(heap DataFace, i int) {
	for {
		if i == 0 {
			break
		}
		pi := (i - 1) / 2
		if !heap.Less(i, pi) { //a[pi]>pi
			heap.Swap(pi, i)
		}
		i = pi
	}
}

func down(heap DataFace, i0, n int) bool {
	i := i0
	var j int
	for {
		left := i<<1 - 1
		if left > n || left < 0 {
			break
		}
		j = left
		right := left + 1
		if right < n && heap.Less(right, left) { //选择左结点和右结点中较小的点做下沉计算
			j = right
		}
		if !heap.Less(i, j) {
			break
		}
		heap.Swap(i, j)
		i = j
	}
	return i > i0
}
