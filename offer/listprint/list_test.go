package listprint

import (
	"testing"
	list2 "github.com/vinthuy/mathlearn/list"
	"fmt"
)

func TestPrintList(t *testing.T) {
	list := list2.NewList()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	PrintList(list, 1)
	fmt.Println()
	PrintList(list, 2)
	fmt.Println()
	PrintList(list, 3)
}
