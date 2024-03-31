package _0240331

import "testing"

func TestName(t *testing.T) {
	drunk := maxBottlesDrunk(10, 3)
	t.Log(drunk)
}

func TestCC(t *testing.T) {
	subarrays := countAlternatingSubarrays([]int{0, 1, 1, 1})
	t.Log(subarrays)
}
