package array

import "testing"

func TestSwap(t *testing.T) {
	a := []int{1, 3, 4, 5}
	Swap(a, 2, 3)
	t.Log(a)
}
