package _0240303

import "testing"

func TestCountSubmatrices(t *testing.T) {
	submatrices := countSubmatrices([][]int{{7, 6, 3}, {6, 6, 1}}, 18)
	t.Log(submatrices)
}
