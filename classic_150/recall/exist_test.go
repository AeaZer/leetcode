package recall

import (
	"testing"
)

func TestExist(t *testing.T) {
	b := exist([][]byte{{'a', 'a'}}, "aaa")
	t.Log(b)
}
