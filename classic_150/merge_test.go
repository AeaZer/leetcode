package classic_150

import (
	"testing"
)

func TestMerge(t *testing.T) {
	res := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {8, 9}, {11, 14}})
	t.Log(res)
}
