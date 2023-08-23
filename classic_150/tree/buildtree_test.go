package tree

import (
	"testing"
)

func TestBuildTreePostorder(t *testing.T) {
	node := buildTreePostorder([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	t.Log(node)
}
