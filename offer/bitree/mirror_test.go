package bitree

import (
	"testing"
	"github.com/vinthuy/mathlearn/tree"
)

func TestMirror(t *testing.T) {
	bi1 := tree.BuildBiTreeFromArray(1, 2, 3, 4, 5)
	bi1.BreadthSearchForWhile(func(v interface{}) {
		t.Log(v)
	})
	Mirror(bi1)
	bi1.BreadthSearchForWhile(func(v interface{}) {
		t.Log(v)
	})
}
