package simulations

import "testing"

func TestCountHillValley(t *testing.T) {
	values := countHillValley([]int{6, 6, 5, 5, 4, 1})
	t.Log(values)
}
