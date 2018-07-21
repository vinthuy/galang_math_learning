package pack

import "testing"

func TestChangeA(t *testing.T) {
	a := A{"1"}
	changeA(a)
	t.Log(a)

	changeAByP(&a)
	t.Log(a)
}
