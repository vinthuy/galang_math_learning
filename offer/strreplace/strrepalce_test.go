package strreplace

import (
	"testing"
)

func TestReplaceBlank(t *testing.T) {
	t.Log(ReplaceBlank("1hello w old", ' ', "20%"))
}
