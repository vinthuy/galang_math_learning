package offer_singleton

import "testing"

func TestSingleton1(t *testing.T) {
	t.Log(SingletonObj.Name)
}

func TestGetSingleton(t *testing.T) {
	t.Log(GetSingleton().Name)
}
