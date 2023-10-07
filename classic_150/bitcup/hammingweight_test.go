package bitcup

import "testing"

func TestHammingWeight(t *testing.T) {
	weight := hammingWeight(00000000000000000000000000001011)
	t.Log(weight)
}
