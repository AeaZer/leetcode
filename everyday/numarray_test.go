package everyday

import "testing"

func TestNumArray(t *testing.T) {
	n := Constructor([]int{1, 3, 5})
	sumRange := n.SumRange(0, 2)
	t.Log(sumRange)
	n.Update(1, 2)
	sumRange = n.SumRange(0, 2)
	t.Log(sumRange)
}
