package listprint

import (
	"testing"
	"github.com/vinthuy/mathlearn/list"
)

func TestLinklist_k(t *testing.T) {
	lst := list.NewList()
	lst.AddFromArray(1, 2, 3, 4, 5)
	t.Log(Linklist_k(lst, 5))
	t.Log(Linklist_k(lst, 6))
}
